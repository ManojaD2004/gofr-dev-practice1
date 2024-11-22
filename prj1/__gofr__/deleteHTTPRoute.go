package __gofr__

import (
	"encoding/json"
	"fmt"
	"gofr.dev/pkg/gofr"
	"strings"
)

func DeleteHTTPRoute(ctx *gofr.Context) (interface{}, error) {
	newRoute := CreateRouteType{}
	ctx.Bind(&newRoute)
	retType := ReturnType{}
	ctx.File.ChDir("./route")
	funcName := toUnderscore(newRoute.RouteName[1:]) + newRoute.Method + "Handler"
	fileName := capWord(toUnderscore(newRoute.RouteName[1:])) + newRoute.Method + ".go"
	f, _ := ctx.File.Open(fileName)
	if f == nil {
		retType.IsDone = false
		retType.Message = "Route does not already exist!"
		ctx.Logger.Info(retType.Message)
		return retType, nil
	}
	f.Close()
	ctx.File.Remove(fileName)
	ctx.File.ChDir("..")
	f, _ = ctx.File.Open("main.go")
	tempS := ""
	reader1, _ := f.ReadAll()
	for reader1.Next() {
		var b string
		reader1.Scan(&b)
		tempS += b + "\n"
	}
	newStr := strings.Replace(tempS, "app."+newRoute.Method+"(\""+newRoute.RouteName+"\", r."+capWord(funcName)+")\n", "", 1)
	f.Close()
	f, _ = ctx.File.Create("main.go")
	n, _ := f.Write([]byte(newStr))
	f.Close()
	fmt.Println("Total bytes written: ", n)
	retType.Message = newRoute.RouteName + " route of " + newRoute.Method + " deleted, and file " + fileName + " deleted!"
	retType.IsDone = true
	ctx.File.ChDir("./__gofr__")
	f, err1 := ctx.File.Open("metadata.json")
	if err1 != nil {
		ctx.Logger.Info("Error opening JSON Object")
		retType.Message = "Error opening JSON Object"
		retType.IsDone = false
		return retType, nil
	}
	read, _ := f.ReadAll()
	mdt := MetaDataType{}
	for read.Next() {
		read.Scan(&mdt)
	}
	delete(mdt.Routes, newRoute.RouteName+"_"+newRoute.Method)
	s1, err := json.Marshal(mdt)
	if err != nil {
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
