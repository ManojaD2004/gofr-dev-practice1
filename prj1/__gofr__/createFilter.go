package __gofr__

import (
	// "encoding/json"
	// "strings"
	"encoding/json"
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
	// fmt.Printf("%v\n", modName)
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
	// fmt.Println(s)
	// fmt.Println(newFilter.Filter)
	ctx.File.ChDir("./filter")
	fileName := toUnderscore(newFilter.FilterName) + ".go"
	f1, _ := ctx.File.Open(fileName)
	if f1 != nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Filter already exist!"
		ctx.Logger.Info(retType.Message)
		return retType, nil
	}
	f, _ = ctx.File.Create(fileName)
	n, _ := f.Write([]byte(s))
	f.Close()
	fmt.Println("Total bytes written: ", n)
	ctx.File.ChDir("..")
	ctx.File.ChDir("./__gofr__")
	f, err1 = ctx.File.Open("metadata.json")
	if err1 != nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Error opening JSON Object"
		ctx.Logger.Info(retType.Message)
		return retType, nil
	}
	s = ""
	read, _ := f.ReadAll()
	mdt := MetaDataType{}
	for read.Next() {
		read.Scan(&mdt)
	}
	mdt.Filter[newFilter.FilterName] = newFilter.Filter
	s1, err := json.Marshal(mdt)
	if err != nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Error converting JSON Object"
		ctx.Logger.Info(retType.Message)
		return retType, nil
	}
	f, _ = ctx.File.Create("metadata.json")
	f.Write(s1)
	f.Close()
	ctx.File.ChDir("..")
	retType.IsDone = true
	retType.Message = "Filter of " + newFilter.FilterName + " created, of type " + newFilter.DataType + "!"
	ctx.Logger.Info(retType.Message)
	return retType, nil
}
