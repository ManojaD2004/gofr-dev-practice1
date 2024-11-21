package main

import (
	__gofr__ "github.com/ManojaD2004/__gofr__"
	r "github.com/ManojaD2004/route"
	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()

	// register routes
	app.POST("/user", r.UserHandler)
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello World", nil
	})
	app.POST("/create-route", __gofr__.CreateRoute)
	// it can be over-ridden through configs
	app.Run()
}
