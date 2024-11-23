package __gofr__

import (
	// "encoding/json"
	// "strings"
	"fmt"

	"gofr.dev/pkg/gofr"
)

func CreateFilter(ctx *gofr.Context) (interface{}, error) {
	newFilter := NewFilterType{}
	ctx.Bind(&newFilter)
	retType := ReturnType{}
	f, _ := ctx.File.Open("./go.mod")
	reader1, err1 := f.ReadAll()
	if err1 != nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Cannot open ./go.mod file!"
		return retType, nil
	}
	s := ""
	reader1.Next()
	var b string
	reader1.Scan(&b)
	modName := b[7:]
	f.Close()
	funcName := capWord(toUnderscore(newFilter.FilterName))
	fmt.Printf("%v\n", modName)
	s = "package filter\n\n" + "import (" + "\n\tt " + `"` + modName + `/types"` + "\n)\n\n"
	s = s + "func " + capWord(funcName) + " " + "(oldD, newD *[]t." + capWord(toUnderscore(newFilter.DataType)) + "Type" + ")" + "{\n"
	s = s + "\tfor i := 0; i < len(*oldD); i++ {\n"
	s = s + "\t\tif !(*oldD)[i].Validate() {\n" + "\t\t\tcontinue" + "\n\t\t}\n"
	s = s + "\t\ta := true\n"
	s = s + "\t\td := (*oldD)[i]\n"
	recurFilter(&newFilter.Filter, &s, 2, "")
	s = s + "\t\tif a {\n" + "\t\t\t*newD = append(*newD, d)" + "\n\t\t}\n"
	s = s + "\t\t// Your extra conversion logic goes here\n"
	s = s + "\t}\n"
	s = s + "}\n"
	fmt.Println(s)
	fmt.Println(newFilter.Filter)
	// fileName := toUnderscore(newFilter.ConversionName) + ".go"
	// f1, _ := ctx.File.Open(fileName)
	return retType, nil
}
