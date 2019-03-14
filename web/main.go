package main

import (
	// "net/http"
	// "log"
	// "fmt"
	// "github.com/gorilla/mux"
	// "github.com/urfave/negroni"
	// "github.com/ghoiufyia/WxApp.kindle/web/modules"
	_ "github.com/ghoiufyia/WxApp.kindle/web/routes"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
)

func main() {
	// port := "6767"
	// log.Println("Starting HTTP service at " + port)
	// fmt.Printf("Starting %v\n", appName)
	
	// service := modules.NewService("app",16)

	// r := mux.NewRouter()
	// service.RegisterRoutes(r,"/v1")
	// http.Handle("/", r)
	
    // err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

    // if err != nil {
    //     log.Println("An error occured starting HTTP listener at port " + port)
    //     log.Println("Error: " + err.Error())
    // }
	
	// dogo.Log.Info("%s,type is %s","file","json")
	
	//初始化配置文件
	cfg := dogo.NewConfig("")
	//new DB
	db,err := database.NewDatabase(cfg)
	if err != nil {
		dogo.Log.Info("init db failed")
	}
	//newApp
	doApp := dogo.NewApp(cfg,db)
	//运行
	doApp.Run();
}