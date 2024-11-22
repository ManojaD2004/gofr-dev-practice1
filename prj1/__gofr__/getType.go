package __gofr__

import (
	"gofr.dev/pkg/gofr"
)

func GetTypeRoute(ctx *gofr.Context) (interface{}, error) {
	getType := GetType{}
	retType := ReturnType{}
	ctx.Bind(&getType)
	ctx.File.ChDir("./types")
	fileName := getType.TypeName + ".go"
	f, _ := ctx.File.Open(fileName)
	if f == nil {
		ctx.Logger.Info("Type/File does not exist")
		retType.Message = "type/file does not exist"
		retType.IsDone = false
		return retType, nil
	}
	f.Close()
	ctx.File.ChDir("..")
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
	retType.Message = mdt.Types[getType.TypeName]
	retType.IsDone = true
	ctx.File.ChDir("..")
	return retType, nil
}
