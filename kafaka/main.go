package main

import (
	"fmt"
    "io"
	"os"
	"gofr.dev/pkg/gofr"
	"mime/multipart"
	"gofr-server/kafaka/producer"
	"gofr-server/kafaka/consumer"
)

type CSVRequest struct {
	TableName string               `form:"tablename"`
	FileName  string               `form:"filename"`
	CSVFile   *multipart.FileHeader `file:"csvfile"` 
}

func main() {

	app := gofr.New()

	app.POST("/convertcsv", func(c *gofr.Context) (interface{}, error) {
		var request CSVRequest

		if err := c.Request.Bind(&request); err != nil {
			return nil, fmt.Errorf("could not bind form data: %v", err)
		}

		if request.CSVFile == nil {
			return nil, fmt.Errorf("no file uploaded with the field name 'csvfile'")
		}

		file, err := request.CSVFile.Open()
		if err != nil {
			return nil, fmt.Errorf("could not open uploaded CSV file: %v", err)
		}
		defer file.Close()

		tempFilePath := fmt.Sprintf("./producer/%s", request.FileName)
		tempFile, err := os.Create(tempFilePath)
		if err != nil {
			return nil, fmt.Errorf("could not create temporary file: %v", err)
		}
		defer tempFile.Close()

		_, err = io.Copy(tempFile, file)
		if err != nil {
			return nil, fmt.Errorf("could not write to temporary file: %v", err)
		}
		// err = processCSV(tempFile.Name())
		// if err != nil {
		// 	return nil, fmt.Errorf("could not process CSV file: %v", err)
		// }
       producer.Producer(request.TableName,request.FileName)
	   consumer.Consumer()
		return fmt.Sprintf("Successfully processed CSV file for table '%s'", request.TableName), nil
	})
	app.Run()
}

// func processCSV(filePath string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return fmt.Errorf("could not open CSV file: %v", err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		return fmt.Errorf("could not read CSV file: %v", err)
// 	}

// 	for _, record := range records {
// 		fmt.Println("Record:", record)
// 	}

// 	return nil
// }
