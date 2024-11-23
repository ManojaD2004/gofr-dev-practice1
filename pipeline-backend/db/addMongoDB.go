package db

import (
	"context"
	"log"
	pipe "github.com/ManojaD2004/pipelines"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddMongoDB(dbName string, conncString string) {
	db, err := mongo.Connect(pipe.C1, options.Client().ApplyURI(conncString))
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = db.Ping(pipe.C1, nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	log.Println("Successfully Connected MongoDB")
	pipe.C1 = context.WithValue(pipe.C1, dbName, db)
}
