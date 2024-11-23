package main

import (
	// "context"
	// "fmt"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	d "github.com/ManojaD2004/db"
	p "github.com/ManojaD2004/pipelines"
	_ "github.com/lib/pq"
	"gofr.dev/pkg/gofr"
	// "github.com/jmoiron/sqlx"
	// "gofr.dev/pkg/gofr"
	// "net/http"
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
	db1 := d.AddPSQL("user=postgres dbname=postgres port=5432 password=pass host=localhost sslmode=disable")
	db2 := d.AddMySQL("root:pass@tcp(127.0.0.1:3306)/mysql")
	stdReport := StudentReport{}
	stdReports := []StudentReport{}
	stdRecs := []Student{}
	f := p.AddPipeline(func() {
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
	}, nil)
	f1 := p.AddPipeline(func() {
		month := time.Now().Month().String()
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
	}, f)
	f2 := p.AddPipeline(func() {
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
		db2.Exec("DELETE FROM student_report;")
		fmt.Println("Both the job have completed, cleaning up the records!!")
	}, f1)
	f3 := p.AddPipeline(func() {
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
	}, f2)
	// db1.Close()
	// db2.Close()
	app := gofr.New()
	app.POST("/try", func(c *gofr.Context) (interface{}, error) {
		f3()
		return "done", nil
	})
	app.Run()
}
