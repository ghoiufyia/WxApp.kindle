package main

import (
	// "net/http"
	// "log"
	// "fmt"
	// "github.com/gorilla/mux"
	// "github.com/urfave/negroni"
	// "github.com/ghoiufyia/WxApp.kindle/web/modules"
	"github.com/ghoiufyia/WxApp.kindle/web/routes"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"

)

func main() {
	// 初始化配置文件
	cfg := dogo.NewConfig("")
	// new DB
	// db,err := dogo.NewDB(cfg.Database)
	// if err != nil {
	// 	dogo.Log.Info("init db failed")
	// }

	// // 设置DB连接池
	// dogo.SetDB(db)
	// new server
	server,err := dogo.NewServer(cfg.Server)
	if err != nil {
		dogo.Log.Info("init server failed")
	}
	// newApp
	doApp := dogo.NewApp(cfg)
	// 注册server
	doApp.RegisterServer(server)
	// 初始化handle
	handler := routes.Init()
	doApp.RegisterHandler(handler)
	//运行
	doApp.Run();
}