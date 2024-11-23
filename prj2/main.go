package main

import (
	"encoding/json"
	"errors"
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
	app.AddPubSub()
	// it can be over-ridden through configs
	app.Subscribe("order-logs", func(c *gofr.Context) error {
		var orderStatus struct {
			OrderId string `json:"orderId"`
			Status  string `json:"status"`
		}
		err := c.Bind(&orderStatus)
		if err != nil {
			c.Logger.Error(err)
			return nil
		}
		fmt.Println("I am in kafka!")
		c.Logger.Info("Received order ", orderStatus)

		return nil
	})
	app.POST("/publish-order", order)
	app.Run()
}

func order(ctx *gofr.Context) (interface{}, error) {
	type orderStatus struct {
		OrderId string `json:"orderId"`
		Status  string `json:"status"`
	}

	var data orderStatus
	b := t.User1{UserName: "manu", Id: 7}
	fmt.Println(b.Validate())
	if !b.Validate() {
		return nil, errors.New("wrong Data Valid")
	}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}
	data.OrderId = "202"
	msg, _ := json.Marshal(data)
	err = ctx.GetPublisher().Publish(ctx, "order-logs", msg)
	if err != nil {
		return nil, err
	}
	err = ctx.GetPublisher().Publish(ctx, "dummy", msg)
	if err != nil {
		return nil, err
	}

	return "Published", nil
}
