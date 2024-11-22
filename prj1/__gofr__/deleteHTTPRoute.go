package __gofr__

import (
	"encoding/json"
	"fmt"
	"gofr.dev/pkg/gofr"
	"strings"
)

func DeleteHTTPRoute(ctx *gofr.Context) (interface{}, error) {
	delRoute := DeleteRouteType{}
	ctx.Bind(&delRoute)
	retType := ReturnType{}
	ctx.File.ChDir("./route")
	funcName := toUnderscore(delRoute.RouteName[1:]) + delRoute.Method + "Handler"
	fileName := capWord(toUnderscore(delRoute.RouteName[1:])) + delRoute.Method + ".go"
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
	newStr := strings.Replace(tempS, "app."+delRoute.Method+"(\""+delRoute.RouteName+"\", r."+capWord(funcName)+")\n\t", "", 1)
	f.Close()
	f, _ = ctx.File.Create("main.go")
	n, _ := f.Write([]byte(newStr))
	f.Close()
	fmt.Println("Total bytes written: ", n)
	retType.Message = delRoute.RouteName + " route of " + delRoute.Method + " deleted, and file " + fileName + " deleted!"
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
	delete(mdt.Routes, delRoute.RouteName+"_"+delRoute.Method)
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
