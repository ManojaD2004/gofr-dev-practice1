package __gofr__

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)

func CreateType(ctx *gofr.Context) (interface{}, error) {
	newType := NewType{}
	retType := ReturnNewType{}
	ctx.Bind(&newType)
	reqType, ok := newType.TypeBody.(map[string]interface{})
	if !ok {
		retType.Message = "invalid request body type"
		retType.IsCreated = false
		return reqType, fmt.Errorf("invalid request body type")
	}
	s := ""
	recur(&reqType, &s, 1)
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
	return retType, nil
}
