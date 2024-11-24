package __gofr__

import (
	// "encoding/json"
	// "strings"
	"encoding/json"
	"gofr.dev/pkg/gofr"
)

func DeleteFilter(ctx *gofr.Context) (interface{}, error) {
	newFilter := DeleteFilterType{}
	ctx.Bind(&newFilter)
	retType := ReturnType{}
	ctx.File.ChDir("./filter")
	fileName := toUnderscore(newFilter.FilterName) + ".go"
	f1, _ := ctx.File.Open(fileName)
	if f1 == nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Filter does not already exist!"
		ctx.Logger.Info(retType.Message)
		return retType, nil
	}
	f1.Close()
	ctx.File.Remove(fileName)
	ctx.File.ChDir("..")
	ctx.File.ChDir("./__gofr__")
	f, err1 := ctx.File.Open("metadata.json")
	if err1 != nil {
		ctx.File.ChDir("..")
		retType.IsDone = false
		retType.Message = "Error opening JSON Object"
		ctx.Logger.Info(retType.Message)
		return retType, nil
	}
	read, _ := f.ReadAll()
	mdt := MetaDataType{}
	for read.Next() {
		read.Scan(&mdt)
	}
	delete(mdt.Filter, newFilter.FilterName)
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
	retType.Message = "Filter of " + newFilter.FilterName + " deleted, of type " + newFilter.DataType + "!"
	ctx.Logger.Info(retType.Message)
	return retType, nil
}
