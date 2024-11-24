package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	d "github.com/ManojaD2004/db"
	"github.com/ManojaD2004/googledrive"
	p "github.com/ManojaD2004/pipelines"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr"
)

type Student struct {
	Name   string  `db:"name"`
	Marks  float64 `db:"marks"`
	Rollno int     `db:"rollno"`
}

type StudentReport struct {
	Month    string  `db:"month"`
	Category string  `db:"category"`
	Name     string  `db:"name"`
	Marks    float64 `db:"marks"`
	Rollno   int     `db:"rollno"`
}

func main() {
	p.C1 = context.Background()
	stdReport := StudentReport{}
	stdReports := []StudentReport{}
	stdRecs := []Student{}
	d.AddPSQL("psql-db1", "user=postgres dbname=postgres port=5432 password=pass host=localhost sslmode=disable")
	d.AddMySQL("mysql-db2", "root:pass@tcp(127.0.0.1:3306)/mysql")
	d.AddMongoDB("mongodb-db3", "mongodb://localhost:27017")
	d.AddRedis("redis-db4", "localhost:6379")
	d.AddCassandra("cassandra-db5", "localhost:9042")
	initFun := p.AddPipeline(nil, func(c1 context.Context) {
		fmt.Println("Init Func!")
	})
	f := p.AddPipeline(initFun, func(c1 context.Context) {
		db1 := p.GetPSQLDB(c1, "psql-db1")
		rows, nil := db1.Queryx("SELECT * FROM students;")
		for rows.Next() {
			stdRec := Student{}
			err := rows.StructScan(&stdRec)
			if err != nil {
				log.Fatal(err)
			}
			stdRecs = append(stdRecs, stdRec)
			fmt.Println(stdRec)
		}

		fmt.Println("Job 1 done!")
	})
	f1 := p.AddPipeline(f, func(c1 context.Context) {
		month := time.Now().Month().String()
		db2 := p.GetMySQLDB(c1, "mysql-db2")
		db2.Exec("DELETE FROM student_report;")
		stdReport.Month = month
		avgMark := 0.0
		for i := 0; i < len(stdRecs); i++ {
			avgMark += stdRecs[i].Marks
		}
		stdReport.Marks = avgMark / float64(len(stdRecs))
		stdReport.Category = "avg marks"
		stdReport.Rollno = 0
		stdReport.Name = "all"
		_, err := db2.Exec("INSERT INTO student_report (month, category, name, marks, rollno) VALUES(?, ?, ?, ?, ?);", stdReport.Month, stdReport.Category, stdReport.Name, stdReport.Marks, stdReport.Rollno)
		if err != nil {
			log.Fatal(err)
		}
		maxMark := 0
		minMark := 0
		for i := 0; i < len(stdRecs); i++ {
			if stdRecs[maxMark].Marks < stdRecs[i].Marks {
				maxMark = i
			}
			if stdRecs[minMark].Marks > stdRecs[i].Marks {
				minMark = i
			}
		}
		stdReport.Marks = stdRecs[maxMark].Marks
		stdReport.Category = "max marks"
		stdReport.Rollno = stdRecs[maxMark].Rollno
		stdReport.Name = stdRecs[maxMark].Name
		_, err = db2.Exec("INSERT INTO student_report (month, category, name, marks, rollno) VALUES(?, ?, ?, ?, ?);", stdReport.Month, stdReport.Category, stdReport.Name, stdReport.Marks, stdReport.Rollno)
		if err != nil {
			log.Fatal(err)
		}
		stdReport.Marks = stdRecs[minMark].Marks
		stdReport.Category = "min marks"
		stdReport.Rollno = stdRecs[minMark].Rollno
		stdReport.Name = stdRecs[minMark].Name
		_, err = db2.Exec("INSERT INTO student_report (month, category, name, marks, rollno) VALUES(?, ?, ?, ?, ?);", stdReport.Month, stdReport.Category, stdReport.Name, stdReport.Marks, stdReport.Rollno)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Job 2 Done!")
	})
	f2 := p.AddPipeline(f1, func(c1 context.Context) {
		db2 := p.GetMySQLDB(c1, "mysql-db2")
		rows, nil := db2.Query("SELECT * FROM student_report;")
		for rows.Next() {
			stdRec := StudentReport{}
			err := rows.Scan(&stdRec.Month, &stdRec.Category, &stdRec.Name, &stdRec.Marks, &stdRec.Rollno)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(stdRec)
			stdReports = append(stdReports, stdRec)
		}
		fmt.Println("Both the job have completed, cleaning up the records!!")
	})
	f3 := p.AddPipeline(f2, func(c1 context.Context) {
		file, err := os.Create("student_reports.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write([]string{"Month", "Category", "Name", "Marks", "Rollno"})

		for _, st := range stdReports {
			record := []string{
				st.Month,
				st.Category,
				st.Name,
				strconv.FormatFloat(st.Marks, 'f', 2, 64),
				strconv.Itoa(st.Rollno),
			}
			err := writer.Write(record)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("Job 3 done, writing to .CSV file!")
	})
	f3_2 := p.AddPipeline(f3, func(c1 context.Context) {
		db3 := p.GetMongoDB(c1, "mongodb-db3")
		stdCollections := db3.Database("mongo-db1").Collection("student_report")
		delResult, err := stdCollections.DeleteMany(c1, bson.M{})
		if err != nil {
			log.Fatal("Error cleaning collections:", err)
		}
		fmt.Println("Deleted all docs in collections:", delResult.DeletedCount)
		for i := 0; i < len(stdReports); i++ {
			stdB := bson.M{"name": stdReports[i].Name, "category": stdReports[i].Category, "marks": stdReports[i].Marks, "month": stdReports[i].Month, "rollno": stdReports[i].Rollno}
			insertResult, err := stdCollections.InsertOne(c1, stdB)
			if err != nil {
				log.Fatal("Error inserting document:", err)
			}
			fmt.Println("Inserted document with ID:", insertResult.InsertedID)
			fmt.Println(stdB)
		}
		fmt.Println("Job 3_2 done, inserting to mongodb!")
	})
	f3_3 := p.AddPipeline(f3_2, func(c1 context.Context) {
		db3 := p.GetRedis(c1, "redis-db4")
		for i := 0; i < len(stdReports); i++ {
			jsonData, _ := json.Marshal(stdReports[i])
			err := db3.Set(c1, stdReports[i].Category, string(jsonData), 0).Err()
			if err != nil {
				log.Fatal("Failed to set key", err)
			}
		}
		stdReports = []StudentReport{}
		fmt.Println("Job 3_3 done, inserting to redis!")
	})
	f3_4 := p.AddPipeline(f3_3, func(c1 context.Context) {
		db := p.GetCassandra(c1, "cassandra-db5")
		err := db.Query("TRUNCATE students;").Exec()
		if err != nil {
			log.Fatal("Error cleaning data, ", err)
		}
		for i := 0; i < len(stdRecs); i++ {
			err = db.Query("INSERT INTO students (name, rollno, marks) VALUES (?, ?, ?);", stdRecs[i].Name, stdRecs[i].Rollno, float32(stdRecs[i].Marks)).Exec()
			if err != nil {
				log.Fatal("Error inserting data, ", err)
			}
		}
		fmt.Println("Done Writing all records to Cassandra DB!")
	})
	f4 := p.AddPipeline(f3_4, func(c1 context.Context) {
		// db1 := p.GetPSQLDB(c1, "psql-db1")
		// db2 := p.GetMySQLDB(c1, "mysql-db2")
		// db1.Close()
		// db2.Close()
		googledrive.Googledrive("credentials.json", "student_reports.csv")
		fmt.Println("Job 4 done, uploading to Google Drive, uncommented the above line!")
	})
	app := gofr.New()
	app.POST("/pipeline/trigger", func(c *gofr.Context) (interface{}, error) {
		f4(p.C1)
		return "done", nil
	})
	app.Run()
}
