package consumer

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/segmentio/kafka-go"
	"gofr-server/producer"
)

func Consumer() {
	ctx := context.Background()
	connectionString := "postgresql://postgres.ejspmxevlxzxowxttbmy:Vilasbhai94!@aws-0-ap-south-1.pooler.supabase.com:6543/postgres"

	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer pool.Close()

	_, err = pool.Exec(ctx, producer.GeneratedQuery)
	if err != nil {
		log.Printf("Failed to create table: %v", err)
	} else {
		fmt.Println("Created table successfully.")
	}
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "consumer-group",
	})
	defer kafkaReader.Close()

	noMessageTimeout := 10 * time.Second

	for {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, noMessageTimeout)
		defer cancel()

		msg, err := kafkaReader.ReadMessage(ctxWithTimeout)
		if err != nil {
			if ctxWithTimeout.Err() == context.DeadlineExceeded {
				fmt.Println("No new messages. Exiting consumer loop.")
				break
			}
			log.Printf("Error reading message: %v", err)
			continue
		}
		msgString := string(msg.Value)
		fmt.Printf("Received message: %s\n", msgString)
		values := parseMessage(msgString)
		if values == nil {
			log.Printf("Skipping invalid message: %s", msgString)
			continue
		}

		_, err = pool.Exec(ctx, producer.GeneratedInsertQuery, values...)
		if err != nil {
			log.Printf("Failed to insert data: %v", err)
		} else {
			fmt.Println("Data inserted successfully.")
		}
	}
}

func parseMessage(msg string) []interface{} {
	parts := strings.Split(msg, ", ")
	values := []interface{}{}

	for _, part := range parts {
		keyValue := strings.SplitN(part, ": ", 2)
		if len(keyValue) == 2 {
			values = append(values, keyValue[1])
		} else {
			log.Printf("Invalid part format: %s", part)
			return nil
		}
	}
	return values
}

// 	err := pool.Exec(ctx, query, firstName, lastName)
// 	if err != nil {
// 	log.Printf("Failed to insert data into DB: %v", err)
// } else {
// 		fmt.Println("Inserted a new tuple in table table_name2 successfully :)")
// 		fmt.Printf("FirstName: %s, LastName: %s\n", firstName, lastName)
// 	}
