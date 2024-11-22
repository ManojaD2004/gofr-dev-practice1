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
	app.DELETE("/user2", r.User2DELETEHandler)
	// Dummy route
	app.GET("/greet", r.GreetRouteGET)
	// HTTP Routes Handler
	app.POST("/.__gofr__/create-route", __gofr__.CreateHTTPRoute)
	app.POST("/.__gofr__/update-route", __gofr__.CreateHTTPRoute)
	app.POST("/.__gofr__/delete-route", __gofr__.DeleteHTTPRoute)
	app.POST("/.__gofr__/get-route", __gofr__.GetHTTPRoute)
	app.GET("/.__gofr__/get-all-routes", __gofr__.GetAllHTTPRoute)
	// Types Routes Handler
	app.POST("/.__gofr__/create-type", __gofr__.CreateTypeRoute)
	app.POST("/.__gofr__/delete-type", __gofr__.DeleteTypeRoute)
	app.POST("/.__gofr__/update-type", __gofr__.UpdateTypeRoute)
	app.POST("/.__gofr__/get-type", __gofr__.GetTypeRoute)
	app.GET("/.__gofr__/get-all-types", __gofr__.GetAllTypesRoute)
	// it can be over-ridden through configs
	app.Run()
}
