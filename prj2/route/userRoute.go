package route

import (
	"gofr.dev/pkg/gofr"
	t "github.com/ManojaD2004/types"
)

func userHandler (ctx *gofr.Context) (interface{}, error) {
	reqBody := t.UserType{}
	ctx.Bind(&reqBody)
	return "Hello World", nil
}

