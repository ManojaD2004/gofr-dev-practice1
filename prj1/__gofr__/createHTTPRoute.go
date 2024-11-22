package __gofr__

import (
	"encoding/json"
	"fmt"
	"gofr.dev/pkg/gofr"
	"strings"
)

func CreateHTTPRoute(ctx *gofr.Context) (interface{}, error) {
	newRoute := CreateRouteType{}
	ctx.Bind(&newRoute)
	retType := ReturnType{}
	f, _ := ctx.File.Open("./go.mod")
	reader1, err1 := f.ReadAll()
	if err1 != nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Cannot open ./go.mod file!"
		return nil, err1
	}
	s := ""
	reader1.Next()
	var b string
	reader1.Scan(&b)
	modName := b[7:]
	f.Close()
	funcName := toUnderscore(newRoute.RouteName[1:]) + newRoute.Method + "Handler"
	fmt.Printf("%v\n", modName)
	s = "package route\n\n" + "import (\n\t" + `"gofr.dev/pkg/gofr"` + "\n\tt " + `"` + modName + `/types"` + "\n)\n\n"
	s = s + "func " + capWord(funcName) + " " + "(ctx *gofr.Context) (interface{}, error) {\n"
	if newRoute.Method != "GET" || newRoute.ReqBodyType != "" {
		s = s + "\treqBody := t." + capWord(toUnderscore(newRoute.ReqBodyType)) + "Type{}\n"
		s = s + "\tctx.Bind(" + "&reqBody" + ")\n"
	}
	s = s + "\tresBody := t." + capWord(toUnderscore(newRoute.ResBodyType)) + "Type{}\n"
	s = s + "\t// Your code logic goes here\n\t\n"
	s = s + "\treturn resBody, nil\n"
	s = s + "}\n\n"
	ctx.File.ChDir("./route")
	fileName := capWord(toUnderscore(newRoute.RouteName[1:])) + newRoute.Method + ".go"
	f1, _ := ctx.File.Open(fileName)
	if f1 != nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Route already exist!"
		ctx.Logger.Info(retType.Message)
		return retType, nil
	}
	f, _ = ctx.File.Create(fileName)
	n, _ := f.Write([]byte(s))
	f.Close()
	fmt.Println("Total bytes written: ", n)
	ctx.File.ChDir("..")
	f, _ = ctx.File.Open("main.go")
	tempS := ""
	reader1, _ = f.ReadAll()
	for reader1.Next() {
		var b string
		reader1.Scan(&b)
		tempS += b + "\n"
	}
	newStr := strings.Replace(tempS, "// register routes", "// register routes\n\t"+"app."+newRoute.Method+"(\""+newRoute.RouteName+"\", r."+capWord(funcName)+")", 1)
	f.Close()
	f, _ = ctx.File.Create("main.go")
	n, _ = f.Write([]byte(newStr))
	f.Close()
	fmt.Println("Total bytes written: ", n)
	retType.Message = newRoute.RouteName + " route of " + newRoute.Method + " created, and file " + fileName + " created!"
	retType.IsDone = true
	ctx.File.ChDir("./__gofr__")
	f, err1 = ctx.File.Open("metadata.json")
	if err1 != nil {
		ctx.File.ChDir("..")
		ctx.Logger.Info("Error opening JSON Object")
		retType.Message = "Error opening JSON Object"
		retType.IsDone = false
		return retType, nil
	}
	s = ""
	read, _ := f.ReadAll()
	mdt := MetaDataType{}
	for read.Next() {
		read.Scan(&mdt)
	}
	mdt.Routes[newRoute.RouteName+"_"+newRoute.Method] = convertStructToMap(newRoute)
	s1, err := json.Marshal(mdt)
	if err != nil {
		ctx.File.ChDir("..")
		ctx.Logger.Info("Error converting JSON Object")
		retType.Message = "Error converting JSON Object"
		retType.IsDone = false
		return retType, nil
	}
	f, _ = ctx.File.Create("metadata.json")
	f.Write(s1)
	f.Close()
	ctx.File.ChDir("..")
	ctx.Logger.Info(retType.Message)
	return retType, nil
}
