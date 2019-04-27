package routes

import (
	"github.com/ghoiufyia/kindle/app/web/kindle/internal/controllers"
	"github.com/ghoiufyia/kindle/library/net/http/klove"

)

func Init() (*klove.RouteGroup) {
	rg := klove.NewRouteGroup()
	rg.Route("GET","/index/",&controllers.IndexController{},"Index","首页")
	// rg.Route("GET","/session/",&controllers.IndexController{},"Session","首页")
	// rg.Route("GET","/list/",&controllers.IndexController{},"List","首页")
	rg.Route("GET","/json/",&controllers.IndexController{},"Json","接口")
	// rg.Route("GET","/rpc/",&controllers.IndexController{},"Rpc","接口")
	
	//静态路由
	rg.Static("/static/","./static/")


	return rg
}