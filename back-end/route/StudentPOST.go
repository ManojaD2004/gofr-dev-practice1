package route

import (
	"errors"

	t "github.com/ManojaD2004/types"
	"gofr.dev/pkg/gofr"
)

func StudentPOSTHandler(ctx *gofr.Context) (interface{}, error) {
	reqBody := t.StudentsType{}
	ctx.Bind(&reqBody)
	if !reqBody.Validate() {
		return nil, errors.New("invalid Data Format")
	}
	resBody := t.IsDoneType{}
	// Your code logic goes here

	resBody.IsHandled = true
	resBody.Message = "this is a proof of validating your data!"
	if !resBody.Validate() {
		return nil, errors.New("invalid Data Format")
	}
	return resBody, nil
}
