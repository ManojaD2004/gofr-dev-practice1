package __gofr__

import (
	// "encoding/json"
	// "strings"
	"gofr.dev/pkg/gofr"
)

func GetAllFilter(ctx *gofr.Context) (interface{}, error) {
	retType := ReturnType{}
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
	f.Close()
	ctx.File.ChDir("..")
	retType.IsDone = true
	retType.Message = mdt.Filter
	ctx.Logger.Info("Passing all the filters that was created!")
	return mdt.Filter, nil
}
