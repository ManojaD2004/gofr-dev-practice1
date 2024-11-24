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
	app.POST("/user123/v2", r.User123_v2POSTHandler)
	app.POST("/student", r.StudentPOSTHandler)
	// Dummy route
	app.GET("/greet", r.GreetRouteGET)
	// HTTP Routes Handler
	app.POST("/.__gofr__/create-route", __gofr__.CreateHTTPRoute)
	app.POST("/.__gofr__/delete-route", __gofr__.DeleteHTTPRoute)
	app.POST("/.__gofr__/update-route", __gofr__.UpdateHTTPRoute)
	app.POST("/.__gofr__/get-route", __gofr__.GetHTTPRoute)
	app.GET("/.__gofr__/get-all-routes", __gofr__.GetAllHTTPRoute)
	// Types Routes Handler
	app.POST("/.__gofr__/create-type", __gofr__.CreateTypeRoute)
	app.POST("/.__gofr__/delete-type", __gofr__.DeleteTypeRoute)
	app.POST("/.__gofr__/update-type", __gofr__.UpdateTypeRoute)
	app.POST("/.__gofr__/get-type", __gofr__.GetTypeRoute)
	app.GET("/.__gofr__/get-all-types", __gofr__.GetAllTypesRoute)
	// Filter Routes Handler
	app.POST("/.__gofr__/create-filter", __gofr__.CreateFilter)
	app.POST("/.__gofr__/delete-filter", __gofr__.DeleteFilter)
	app.GET("/.__gofr__/get-all-filters", __gofr__.GetAllFilter)
	// it can be over-ridden through configs
	app.Run()
}
