package route

import (
	"fmt"

	t "github.com/ManojaD2004/types"
	"gofr.dev/pkg/gofr"
)

func UserHandler(ctx *gofr.Context) (interface{}, error) {
	reqBody := t.UserType{}
	ctx.Bind(&reqBody)
	fmt.Println(reqBody)
	return "Hello World", nil
}
