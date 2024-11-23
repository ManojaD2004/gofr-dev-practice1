package db

import (
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func AddPSQL(conncString string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", conncString)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil
	} else {
		log.Println("Successfully Connected")
	}
	return db
}
