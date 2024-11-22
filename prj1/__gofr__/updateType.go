package __gofr__

import (
	"encoding/json"
	"fmt"

	"gofr.dev/pkg/gofr"
)

func UpdateTypeRoute(ctx *gofr.Context) (interface{}, error) {
	newType := NewType{}
	retType := ReturnType{}
	ctx.Bind(&newType)
	s := ""
	recur(&newType.TypeBody, &s, 1)
	typeName1 := toUnderscore(newType.TypeName)
	s = "type " + capWord(typeName1) + "Type " + s
	s = "package types\n\n" + s
	ctx.File.ChDir("./types")
	fileName := typeName1 + ".go"
	f, _ := ctx.File.Open(fileName)
	if f == nil {
		ctx.Logger.Info("Type/File does not exist")
		retType.Message = "type/file does not exist"
		retType.IsDone = false
		return retType, nil
	}
	f, _ = ctx.File.Create(fileName)
	n, _ := f.Write([]byte(s))
	f.Close()
	ctx.File.ChDir("..")
	fmt.Println("Total bytes written: ", n)
	retType.Message = newType.TypeName + " type created, and type file " + fileName + " was updated!"
	retType.IsDone = true
	ctx.File.ChDir("./__gofr__")
	f, err := ctx.File.Open("metadata.json")
	if err != nil {
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
	mdt.Types[newType.TypeName] = newType.TypeBody
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
