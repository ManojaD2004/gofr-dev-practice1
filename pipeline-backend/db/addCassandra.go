package db

import (
	"context"
	pipe "github.com/ManojaD2004/pipelines"
	"github.com/gocql/gocql"
	"log"
	"strconv"
	"strings"
)

func AddCassandra(dbName string, conncString string) {
	cluster := gocql.NewCluster(strings.Split(conncString, ":")[0])
	port, _ := strconv.Atoi(strings.Split(conncString, ":")[1])
	cluster.Port = port
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Could not connect to Cassandra", err)
	}
	log.Println("Successfully Connected Cassandra!")
	err = session.Query(`CREATE KEYSPACE IF NOT EXISTS db 
		WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec()
	if err != nil {
		log.Fatal("Failed to create keyspace:", err)
	}
	log.Println("Keyspace created or already exists.")
	cluster.Keyspace = "db"
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Failed to switch keyspace:", err)
	}
	log.Println("Switched to keyspace: db")
	pipe.C1 = context.WithValue(pipe.C1, dbName, session)
}
