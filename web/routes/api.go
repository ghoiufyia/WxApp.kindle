package routes

import (
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"github.com/ghoiufyia/WxApp.kindle/web/controllers"

)

func init() {
	dogo.Router("/api/",&controllers.IndexController{})
}

// var Api = Routes{
// 	Prefix:		"/api/v1",
// 	MyRoutes:	[]Route{
// 		Route{
// 			Name:			"index",
// 			Method:			"GET",
// 			Pattern:		"/index",
// 			HandlerFunc:	handlers.Index,
// 		},
// 	},
// }


// AddRouter("index","GET","/index",&handlers.IndexHandler{})

