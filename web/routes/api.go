package routes

import (
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"github.com/ghoiufyia/WxApp.kindle/web/controllers"

)

func Init() (*dogo.RouteMap) {
	rm := dogo.NewRouteMap()
	rm.Router("首页","GET","/index/",&controllers.IndexController{},"Index")
	rm.Router("首页","GET","/list/",&controllers.IndexController{},"List")
	rm.Router("接口","GET","/json/",&controllers.IndexController{},"Json")
	rm.Router("接口","GET","/rpc/",&controllers.IndexController{},"Rpc")
	return rm
}