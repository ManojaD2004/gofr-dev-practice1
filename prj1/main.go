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
	app.DELETE("/user", r.UserDELETEHandler)
	app.POST("/user/v1/create-users", r.User_v1_create_usersPOSTHandler)
	app.GET("/greet", r.GreetRouteGET)
	app.POST("/create-route", __gofr__.CreateRoute)
	app.POST("/.__gofr__/create-type", __gofr__.CreateTypeRoute)
	app.POST("/.__gofr__/delete-type", __gofr__.DeleteTypeRoute)
	app.POST("/.__gofr__/get-type", __gofr__.GetTypeRoute)
	app.POST("/.__gofr__/update-type", __gofr__.UpdateTypeRoute)
	// it can be over-ridden through configs
	app.Run()
}
