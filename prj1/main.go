package main

import (
	"fmt"
	t "github.com/ManojaD2004/types"
	"gofr.dev/pkg/gofr"
)

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
		for key, value := range reqType {
			switch val := value.(type) {
			case string:
				fmt.Println("This is a string", val)
			case map[string]interface{}:
				fmt.Println("This is a json", val)
			default:
				fmt.Println("Error!")
			}
			fmt.Println("Key: ", key, " Value: ", value)
		}
		return "Hello Tiger!", nil
	})
	// it can be over-ridden through configs
	app.Run()
}
