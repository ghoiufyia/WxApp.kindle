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
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"

)

var (
	DoDb *gorm.DB
)

func main() {
	//初始化配置文件
	cfg := dogo.NewConfig("")
	//new DB
	DoDb,err := dogo.NewDatabase(cfg.Database)
	if err != nil {
		dogo.Log.Info("init db failed")
	}
	server,err := dogo.NewServer(cfg.Server)
	if err != nil {
		dogo.Log.Info("init server failed")
	}
	//newApp
	doApp := dogo.NewApp(cfg)
	// //注册Db
	// DoDb.RegisterDb(db)
	// 注册
	doApp.RegisterServer(server)

	handler := routes.Init()
	doApp.RegisterHandler(handler)
	//运行
	doApp.Run();
}