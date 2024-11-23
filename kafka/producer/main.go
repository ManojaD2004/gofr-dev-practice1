package producer

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/segmentio/kafka-go"
)

var GeneratedQuery string
var GeneratedInsertQuery string
var tableName = "table6"

func Producer() {
	var firstRow []string
	ctx := context.Background()
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	})

	file, err := os.Open("./producer/data.csv")
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

	GeneratedQuery = generateCreateTableQuery(firstRow)
	GeneratedInsertQuery = generateInsertQuery(firstRow)

	for _, record := range records[1:] { 
		message := generateKafkaMessage(firstRow, record)
		err := kafkaWriter.WriteMessages(ctx, kafka.Message{
			Value: []byte(message),
		})
		if err != nil {
			log.Fatalf("Error writing message to Kafka: %v", err)
		}
		fmt.Println("Sent to Kafka:", message)
	}

	defer kafkaWriter.Close()
}

func generateCreateTableQuery(firstRow []string) string {
	columnNames := []string{}
	for _, column := range firstRow {
		columnNames = append(columnNames, fmt.Sprintf("%s VARCHAR(10)", column))
	}
	return fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(columnNames, ", "))
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

func generateKafkaMessage(columns []string, record []string) string {
	var messageParts []string
	for i, value := range record {
		if i < len(columns) { 
			messageParts = append(messageParts, fmt.Sprintf("%s: %s", columns[i], value))
		}
	}
	return strings.Join(messageParts, ", ")
}

