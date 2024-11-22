package __gofr__

import (
	"encoding/json"
	"fmt"
	"gofr.dev/pkg/gofr"
)

func DeleteTypeRoute(ctx *gofr.Context) (interface{}, error) {
	delType := DeleteType{}
	retType := ReturnNewType{}
	ctx.Bind(&delType)
	ctx.File.ChDir("./types")
	fileName := delType.TypeName + ".go"
	f, _ := ctx.File.Open(fileName)
	if f == nil {
		ctx.Logger.Info("Type/File does not exist")
		retType.Message = "type/file does not exist"
		retType.IsDone = false
		return retType, nil
	}
	f.Close()
	ctx.File.Remove(fileName)
	ctx.File.ChDir("..")
	fmt.Println("Type/File deleted")
	retType.Message = delType.TypeName + " type deleted, and type file " + fileName + " deleted!"
	retType.IsDone = true
	ctx.File.ChDir("./__gofr__")
	f, err := ctx.File.Open("metadata.json")
	if err != nil {
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
	delete(mdt.Types, delType.TypeName)
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
