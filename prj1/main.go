package main

import (
	"fmt"
	"reflect"
	"strings"
	r "github.com/ManojaD2004/route"
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

func recur(m *map[string]interface{}, s *string) {
	*s += "struct {\n"
	for key, value := range *m {
		switch val := value.(type) {
		case string:
			*s = *s + "\t" + capitalizeFirstLetter(key) + " " + val + " `json:\"" + key + "\"`\n"
			// fmt.Println(key1, key, val, s)
		case map[string]interface{}:
			// fmt.Println(key1, key, val)
			newType := capitalizeFirstLetter(key)
			*s = *s + "\t" + newType + " "
			recur(&val, s)
			*s = *s + " `json:\"" + key + "\"`\n"
		case []interface{}:
			if len(val) == 1 {
				switch val1 := val[0].(type) {
				case string:
					*s = *s + "\t" + capitalizeFirstLetter(key) + " []" + val1 + " `json:\"" + key + "\"`\n"
					// fmt.Println("Array of string")
				case map[string]interface{}:
					// fmt.Println("Array of object")
					newType := capitalizeFirstLetter(key)
					*s = *s + "\t" + newType + " []"
					recur(&val1, s)
					*s = *s + " `json:\"" + key + "\"`\n"
				default:
					fmt.Println("Error 2!", reflect.TypeOf(val1))
				}
			}
		default:
			fmt.Println("Error 1!", reflect.TypeOf(val))
		}
	}
	*s += "}"
}

func main() {
	// initialise gofr object
	app := gofr.New()
	// register routes
	app.POST("/user", r.UserHandler)
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
		fmt.Println(newType)
		recur(&reqType, &s)
		s = "type " + capitalizeFirstLetter(newType.TypeName) + "Type " + s
		s = "package types\n\n" + s
		ctx.File.ChDir("./types")
		f, _ := ctx.File.Create(smallFirstLetter(newType.TypeName) + ".go")
		n, _ := f.Write([]byte(s))
		fmt.Println(n)
		s = "package route\n\n" + "import (\n\t" + `"gofr.dev/pkg/gofr"` + "\n\tt " + `"github.com/ManojaD2004/types"` + "\n)\n\n"
		s = s + "func " + smallFirstLetter(newType.TypeName) + "Handler " + "(ctx *gofr.Context) (interface{}, error) {\n"
		s = s + "\treqBody := t." + capitalizeFirstLetter(newType.TypeName) + "Type{}\n"
		s = s + "\tctx.Bind(" + "&reqBody" + ")\n"
		s = s + "\treturn \"Hello World\", nil\n"
		s = s + "}\n\n"
		ctx.File.ChDir("../route")
		f, _ = ctx.File.Create(smallFirstLetter(newType.TypeName) + "Route" + ".go")
		n, _ = f.Write([]byte(s))
		
		fmt.Println(n)
		return "Hello Tiger!", nil
	})
	// it can be over-ridden through configs
	app.Run()
}
