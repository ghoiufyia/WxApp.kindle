package routes

import (
	"WxApp.kindle/web/handlers"
)

var Api = Routes{
	Prefix:		"/api/v1",
	MyRoutes:	[]Route{
		Route{
			Name:			"index",
			Method:			"GET",
			Pattern:		"/index",
			HandlerFunc:	handlers.Index,
		},
	},
}


// AddRouter("index","GET","/index",&handlers.IndexHandler{})

