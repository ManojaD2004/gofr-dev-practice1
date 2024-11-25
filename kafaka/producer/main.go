package producer

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
  "time"
	"github.com/segmentio/kafka-go"
)

var GeneratedQuery string
var GeneratedInsertQuery string
var dataTypes []string
var tableName string
func Producer(TableName string,fileName string) {
	tableName = TableName
	var firstRow []string
	ctx := context.Background()
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		BatchSize:     10000,              
		BatchTimeout:  500 * time.Millisecond,
	})
	tempFilePath := fmt.Sprintf("./producer/%s", fileName)
	file, err := os.Open(tempFilePath)
	if err != nil {
		log.Fatalf("Could not open the CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error while reading file %v", err)
	}
	if len(records) > 0 {
		firstRow = records[0]
		log.Println("First row:", firstRow)
	} else {
		log.Println("The CSV file is empty.")
	}

	if len(records) > 1 {
        secondRow := records[1] 
        dataTypes = getDataTypes(secondRow) 
        log.Println("Second row data types:", dataTypes)
    }

    GeneratedQuery = generateCreateTableQueryWithTypes(firstRow, dataTypes) 
    log.Println("Generated CREATE TABLE query:", GeneratedQuery)
    GeneratedInsertQuery = generateInsertQuery(firstRow) 
	log.Println("Generated INSERT TABLE query:", GeneratedQuery)
	j := 1
    for _, record := range records[1:] {
        message := generateKafkaMessage(firstRow, record)
        err := kafkaWriter.WriteMessages(ctx, kafka.Message{
            Value: []byte(message),
        })
        if err != nil {
            log.Fatalf("Error writing message to Kafka: %v", err)
        }
        fmt.Printf("%d) Sent to Kafka: %s\n",j, message)
		j++
    }

    defer kafkaWriter.Close()
}

func getDataTypes(row []string) []string {
    dataTypes := []string{}
    for _, value := range row {
        if isNumeric(value) {
            dataTypes = append(dataTypes, "NUMERIC")
        } else if isBool(value){
			dataTypes = append(dataTypes, "BOOLEAN")
		}else{
				
				dataTypes = append(dataTypes, "VARCHAR(20)")
		}
    }
    return dataTypes
}

func isNumeric(value string) bool {
    _, err := strconv.ParseFloat(value, 64) 
    return err == nil
}
func isBool(value string) bool{
	_,err := strconv.ParseBool(value)
	return err ==nil

}

func generateCreateTableQueryWithTypes(firstRow, dataTypes []string) string {
    columnDefinitions := []string{}
    for i, column := range firstRow {
        columnDefinitions = append(columnDefinitions, fmt.Sprintf("%s %s", column, dataTypes[i]))
    }
    return fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(columnDefinitions, ", "))
}

func generateKafkaMessage(columns []string, record []string) string {
	var messageParts []string
	for i, value := range record {
		if i < len(columns) { 
			messageParts = append(messageParts, fmt.Sprintf("%s: %s", columns[i], value))
		}
	}
	return strings.Join(messageParts, ", ")
}

func generateInsertQuery(firstRow []string) string {
	columnNames := []string{}
	placeholders := []string{}
	for i, column := range firstRow {
		columnNames = append(columnNames, column)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		tableName,
		strings.Join(columnNames, ", "),
		strings.Join(placeholders, ", "),
	)
	return query
}

