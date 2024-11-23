package db

import (
	"context"
	pipe "github.com/ManojaD2004/pipelines"
	"github.com/redis/go-redis/v9"
	"log"
)

func AddRedis(dbName string, conncString string) {
	db := redis.NewClient(&redis.Options{
		Addr:     conncString, // Redis server address
		Password: "",          // No password by default
		DB:       0,           // Default DB
	})
	_, err := db.Ping(pipe.C1).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis", err)
	} else {
		log.Println("Successfully Connected Redis!")
	}
	pipe.C1 = context.WithValue(pipe.C1, dbName, db)
}
