package pipelines

import (
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

var C1 context.Context

func GetPSQLDB(c context.Context, psqlKey string) *sqlx.DB {
	db, ok := c.Value(psqlKey).(*sqlx.DB)
	if !ok {
		log.Fatal("No PSQL database connection found in context")
	}
	return db
}

func GetMySQLDB(c context.Context, mysqlKey string) *sql.DB {
	db, ok := c.Value(mysqlKey).(*sql.DB)
	if !ok {
		log.Fatal("No MySQL database connection found in context")
	}
	return db
}


func GetMongoDB(c context.Context, mysqlKey string) *mongo.Client {
	db, ok := c.Value(mysqlKey).(*mongo.Client)
	if !ok {
		log.Fatal("No MySQL database connection found in context")
	}
	return db
}
