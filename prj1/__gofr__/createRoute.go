package __gofr__

import (
	"fmt"
	"strings"

	"gofr.dev/pkg/gofr"
)

func CreateRoute(ctx *gofr.Context) (interface{}, error) {
	newType := RouteType{}
	ctx.Bind(&newType)
	reqType, ok := newType.ReqBody.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid request body type")
	}
	s := ""
	recur(&reqType, &s, 1)
	s = "type " + capWord(newType.TypeName) + "Type " + s
	s = "package types\n\n" + s
	ctx.File.ChDir("./types")
	f, _ := ctx.File.Create(smallFirstLetter(newType.TypeName) + ".go")
	n, _ := f.Write([]byte(s))
	ctx.File.ChDir("..")
	fmt.Println(n)
	f, _ = ctx.File.Open("./go.mod")
	reader1, err1 := f.ReadAll()
	if err1 != nil {
		return nil, err1
	}
	reader1.Next()
	var b string
	reader1.Scan(&b)
	modName := b[7:]
	fmt.Printf("%v\n", modName)
	s = "package route\n\n" + "import (\n\t" + `"gofr.dev/pkg/gofr"` + "\n\tt " + `"` + modName + `/types"` + "\n)\n\n"
	s = s + "func " + capWord(newType.TypeName) + "Handler " + "(ctx *gofr.Context) (interface{}, error) {\n"
	s = s + "\treqBody := t." + capWord(newType.TypeName) + "Type{}\n"
	s = s + "\tctx.Bind(" + "&reqBody" + ")\n"
	s = s + "\treturn \"Hello World\", nil\n"
	s = s + "}\n\n"
	ctx.File.ChDir("./route")
	f1, _ := ctx.File.Open(smallFirstLetter(newType.TypeName) + "Route" + ".go")
	if f1 == nil {
		ctx.File.ChDir("..")
		f2, _ := ctx.File.Open("main.go")
		tempS := ""
		reader2, _ := f2.ReadAll()
		for reader2.Next() {
			var b string
			reader2.Scan(&b)
			tempS += b + "\n"
		}
		newStr := strings.Replace(tempS, "// register routes", "// register routes\n\t"+`app.POST("/`+smallFirstLetter(newType.TypeName)+`",r.`+capWord(newType.TypeName)+`Handler)`, 1)
		fmt.Println(newStr)
		f2.Close()
		f3, _ := ctx.File.Create("main.go")
		f3.Write([]byte(newStr))
		f3.Close()
		ctx.File.ChDir("./route")
	}
	f, _ = ctx.File.Create(smallFirstLetter(newType.TypeName) + "Route" + ".go")
	n, _ = f.Write([]byte(s))
	f.Close()
	fmt.Println(n)

	return "Hello Tiger!", nil
}
