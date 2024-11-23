package pipelines

import (
	"context"
	"database/sql"
	"log"

	"github.com/gocql/gocql"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
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
		log.Fatal("No MongoDB database connection found in context")
	}
	return db
}


func GetRedis(c context.Context, mysqlKey string) *redis.Client {
	db, ok := c.Value(mysqlKey).(*redis.Client)
	if !ok {
		log.Fatal("No Redis database connection found in context")
	}
	return db
}

func GetCassandra(c context.Context, mysqlKey string) *gocql.Session {
	db, ok := c.Value(mysqlKey).(*gocql.Session)
	if !ok {
		log.Fatal("No Cassandra database connection found in context")
	}
	return db
}
