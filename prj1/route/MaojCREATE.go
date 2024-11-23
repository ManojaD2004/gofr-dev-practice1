package route

import (
	"gofr.dev/pkg/gofr"
	t "github.com/ManojaD2004/types"
)

func MaojCREATEHandler (ctx *gofr.Context) (interface{}, error) {
	reqBody := t.IsDoneType{}
	ctx.Bind(&reqBody)
	resBody := t.User2Type{}
	// Your code logic goes here
	
	return resBody, nil
}

