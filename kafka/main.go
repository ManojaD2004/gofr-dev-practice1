package main

import (
	"fmt"
	"gofr-server/consumer"
	"gofr-server/producer"
)
func main(){
  producer.Producer()
  consumer.Consumer()
	query := producer.GeneratedQuery
	fmt.Println("The generated create query is :")
	fmt.Println(query)
	query2 := producer.GeneratedInsertQuery
	fmt.Println("The generated create query is :")
	fmt.Println(query2)
}