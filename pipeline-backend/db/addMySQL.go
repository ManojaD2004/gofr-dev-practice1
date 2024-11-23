package db

import (
	"context"
	"database/sql"

	pipe "github.com/ManojaD2004/pipelines"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func AddMySQL(dbName string, conncString string) {
	db, err := sql.Open("mysql", conncString)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("Successfully Connected MySQL")
	}
	pipe.C1 = context.WithValue(pipe.C1, dbName, db)
}
