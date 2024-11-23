package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddMySQL(conncString string) *sql.DB {
	db, err := sql.Open("mysql", conncString)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil
	} else {
		log.Println("Successfully Connected")
	}
	return db
}
