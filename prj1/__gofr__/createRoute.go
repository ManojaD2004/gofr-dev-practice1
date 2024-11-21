package __gofr__

import (
	"fmt"
	t "github.com/ManojaD2004/types"
	"gofr.dev/pkg/gofr"
)

func CreateRoute(ctx *gofr.Context) (interface{}, error) {
	newType := t.RouteType{}
	ctx.Bind(&newType)
	reqType, ok := newType.ReqBody.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid request body type, expected map[string]interface{}")
	}
	s := ""
	fmt.Println(newType)
	recur(&reqType, &s, 1)
	s = "type " + capitalizeFirstLetter(newType.TypeName) + "Type " + s
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
	s = "package route\n\n" + "import (\n\t" + `"gofr.dev/pkg/gofr"` + "\n\tt " + `"github.com/ManojaD2004/types"` + "\n)\n\n"
	s = s + "func " + capitalizeFirstLetter(newType.TypeName) + "Handler " + "(ctx *gofr.Context) (interface{}, error) {\n"
	s = s + "\treqBody := t." + capitalizeFirstLetter(newType.TypeName) + "Type{}\n"
	s = s + "\tctx.Bind(" + "&reqBody" + ")\n"
	s = s + "\treturn \"Hello World\", nil\n"
	s = s + "}\n\n"
	ctx.File.ChDir("../route")
	f, _ = ctx.File.Create(smallFirstLetter(newType.TypeName) + "Route" + ".go")
	n, _ = f.Write([]byte(s))

	fmt.Println(n)
	return "Hello Tiger!", nil
}
