package route

import "gofr.dev/pkg/gofr"

func GreetRouteGET(ctx *gofr.Context) (interface{}, error) {
	return "Hello World", nil
}
