package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	t "github.com/ManojaD2004/types"
	"gofr.dev/pkg/gofr"
)

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func smallFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

func recur(m *map[string]interface{}, key1 string, s *string) {
	s1 := "type " + capitalizeFirstLetter(key1) + "Type" + " struct {\n"
	for key, value := range *m {
		switch val := value.(type) {
		case string:
			s1 = s1 + "\t" + capitalizeFirstLetter(key) + " " + val + " `json:\"" + key + "\"`\n"
			// fmt.Println(key1, key, val, s)
		case map[string]interface{}:
			// fmt.Println(key1, key, val)
			newType := capitalizeFirstLetter(key1 + capitalizeFirstLetter(key))
			s1 = s1 + "\t" + newType + " " + newType + "Type" + " `json:\"" + key + "\"`\n"
			recur(&val, newType, s)
		default:
			fmt.Println("Error!")
		}
	}
	s1 += "}\n\n"
	*s = s1 + *s
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
		s := ""
		recur(&reqType, newType.TypeName, &s)
		s = "package types\n\n" + s
		ctx.File.ChDir("./types")
		f, _ := ctx.File.Create(smallFirstLetter(newType.TypeName) + ".go")
		n, _ := f.Write([]byte(s))
		fmt.Println(s, n)
		s = "package route\n\n" + "import (\n\t" + `"gofr.dev/pkg/gofr"` + "\n\tt " + `"github.com/ManojaD2004/types"` + "\n)\n\n"
		s = s + "func " + smallFirstLetter(newType.TypeName) + "Handler " + "(ctx *gofr.Context) (interface{}, error) {\n"
		s = s + "\treqBody := t." + capitalizeFirstLetter(newType.TypeName) + "Type{}\n"
		s = s + "\tctx.Bind(" + "&reqBody" + ")\n"
		s = s + "\treturn \"Hello World\", nil\n"
		s = s + "}\n\n"
		ctx.File.ChDir("../route")
		f, _ = ctx.File.Create(smallFirstLetter(newType.TypeName) + "Route" + ".go")
		n, _ = f.Write([]byte(s))
		fmt.Println(s, n)
		return "Hello Tiger!", nil
	})
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
