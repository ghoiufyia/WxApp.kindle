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

	"github.com/ghoiufyia/WxApp.kindle/web/until/account"

)

func main() {

	acfg := account.NewConfig("./config/account.json");
	dogo.Info(acfg)

	// 初始化配置文件
	cfg := dogo.NewConfig("")
	// 初始化日志
	dogo.InitLog(cfg.Log)
	dogo.Info("测试")
	// 初始化DB
	dogo.InitDatabse(cfg.Database)
	dogo.Info("测试db",dogo.DoDB)
	// new server
	server,err := dogo.NewServer(cfg.Server)
	if err != nil {
		dogo.Info("init server failed")
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