package __gofr__

import (
	"encoding/json"
	"fmt"

	"gofr.dev/pkg/gofr"
)

func CreateType(ctx *gofr.Context) (interface{}, error) {
	newType := NewType{}
	retType := ReturnNewType{}
	ctx.Bind(&newType)
	s := ""
	recur(&newType.TypeBody, &s, 1)
	typeName1 := toUnderscore(newType.TypeName)
	s = "type " + capWord(typeName1) + "Type " + s
	s = "package types\n\n" + s
	ctx.File.ChDir("./types")
	fileName := typeName1 + ".go"
	f, _ := ctx.File.Open(fileName)
	if f != nil {
		retType.Message = "file already exist"
		retType.IsCreated = false
		return retType, nil
	}
	f.Close()
	f, _ = ctx.File.Create(fileName)
	n, _ := f.Write([]byte(s))
	f.Close()
	ctx.File.ChDir("..")
	fmt.Println(n)
	retType.Message = newType.TypeName + " created, and type file " + fileName + "created!"
	retType.IsCreated = true
	ctx.File.ChDir("./")
	f, err := ctx.File.Open("metadata.json")
	if err != nil {
		ctx.Logger.Info("Error opening JSON Object")
	}
	read, _ := f.ReadAll()
	s = ""
	for read.Next() {
		var b string
		read.Scan(&b)
		s += b
	}
	mdt := MetaDataType{}
	json.Unmarshal([]byte(s), &mdt)
	s = ""
	mdt.Types[newType.TypeName] = newType.TypeBody
	s1, err := json.Marshal(mdt)
	if err != nil {
		ctx.Logger.Info("Error converting JSON Object")
	}
	f.Close()
	f, _ = ctx.File.Open("metadata.json")
	f.Write(s1)
	return retType, nil
}
