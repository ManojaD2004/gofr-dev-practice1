package db

import (
	"context"
	pipe "github.com/ManojaD2004/pipelines"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func AddPSQL(dbName string, conncString string) {
	db, err := sqlx.Connect("postgres", conncString)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("Successfully Connected PSQL")
	}
	pipe.C1 = context.WithValue(pipe.C1, dbName, db)
}
