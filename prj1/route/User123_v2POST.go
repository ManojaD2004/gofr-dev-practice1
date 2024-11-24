package route

import (
	"gofr.dev/pkg/gofr"
	t "github.com/ManojaD2004/types"
	"errors"
)

func User123_v2POSTHandler (ctx *gofr.Context) (interface{}, error) {
	reqBody := t.User123Type{}
	ctx.Bind(&reqBody)
	if !reqBody.Validate() {
		return nil, errors.New("invalid Data Format")
	}
	resBody := t.IsDoneType{}
	// Your code logic goes here
	
	if !resBody.Validate() {
		return nil, errors.New("invalid Data Format")
	}
	return resBody, nil
}

