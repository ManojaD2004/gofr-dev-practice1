package route

import (
	"gofr.dev/pkg/gofr"
	t "github.com/ManojaD2004/types"
)

func User2DELETEHandler (ctx *gofr.Context) (interface{}, error) {
	reqBody := t.UserType{}
	ctx.Bind(&reqBody)
	resBody := t.IsDoneType{}
	// Your code logic goes here
	
	return resBody, nil
}

