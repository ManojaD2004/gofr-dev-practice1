package __gofr__

import (
	"gofr.dev/pkg/gofr"
)

func GetHTTPRoute(ctx *gofr.Context) (interface{}, error) {
	getRoute := GetRouteType{}
	ctx.Bind(&getRoute)
	retType := ReturnType{}
	ctx.File.ChDir("./route")
	funcName := toUnderscore(getRoute.RouteName[1:]) + getRoute.Method + "Handler"
	fileName := capWord(toUnderscore(getRoute.RouteName[1:])) + getRoute.Method + ".go"
	f, _ := ctx.File.Open(fileName)
	if f == nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Route does not exist!"
		ctx.Logger.Info(retType.Message)
		return retType, nil
	}
	f.Close()
	ctx.File.ChDir("..")
	ctx.File.ChDir("./__gofr__")
	f, err1 := ctx.File.Open("metadata.json")
	if err1 != nil {
		ctx.File.ChDir("..")
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
	f.Close()
	ctx.File.ChDir("..")
	retType.IsDone = true
	retType.Message = mdt.Routes[getRoute.RouteName+"_"+getRoute.Method]
	ctx.Logger.Info("Sending route " + fileName + " with a function " + funcName)
	return retType, nil
}

func GetAllHTTPRoute(ctx *gofr.Context) (interface{}, error) {
	ctx.File.ChDir("./__gofr__")
	f, err1 := ctx.File.Open("metadata.json")
	if err1 != nil {
		ctx.File.ChDir("..")
		ctx.Logger.Info("Error opening JSON Object")
		retType := ReturnType{}
		retType.Message = "Error opening JSON Object"
		retType.IsDone = false
		return retType, nil
	}
	read, _ := f.ReadAll()
	mdt := MetaDataType{}
	for read.Next() {
		read.Scan(&mdt)
	}
	f.Close()
	ctx.File.ChDir("..")
	return mdt.Routes, nil
}
