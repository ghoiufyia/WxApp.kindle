package routes

import (
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"github.com/ghoiufyia/WxApp.kindle/web/controllers"

)

func Init() (*dogo.RouteMap) {
	rm := dogo.NewRouteMap()
	rm.Router("首页","GET","/index/",&controllers.IndexController{},"Index")
	return rm
}