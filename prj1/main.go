package main

import (
	"fmt"
	t "github.com/ManojaD2004/types"
	"gofr.dev/pkg/gofr"
)

func recur(m *map[string]interface{}, key1 string) {
	for key, value := range *m {
		switch val := value.(type) {
		case string:
			fmt.Println(key1, key, val)
		case map[string]interface{}:
			fmt.Println(key1, key, val)
			recur(&val, key1 + " " + key)
		default:
			fmt.Println("Error!")
		}
		// fmt.Println("Key: ", key, " Value: ", value)
	}
}

func main() {
	// initialise gofr object
	app := gofr.New()
	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello World", nil
	})
	app.POST("/create-route", func(ctx *gofr.Context) (interface{}, error) {
		newType := t.RouteType{}
		ctx.Bind(&newType)
		reqType, ok := newType.ReqBody.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid request body type, expected map[string]interface{}")
		}
		recur(&reqType, newType.TypeName)
		return "Hello Tiger!", nil
	})
	// it can be over-ridden through configs
	app.Run()
}
