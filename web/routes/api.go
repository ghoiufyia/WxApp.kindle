package routes

import (
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"github.com/ghoiufyia/WxApp.kindle/web/controllers"

)

func init() {
	dogo.Router("/api/",&controllers.IndexController{})
}

func Init() {
	
}