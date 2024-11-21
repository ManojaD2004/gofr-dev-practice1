package main

import (
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
		return "Hello Tiger!", nil
	})
	// it can be over-ridden through configs
	app.Run()
}
