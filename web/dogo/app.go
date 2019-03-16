package dogo

import (
	"net/http"
	"fmt"
)

// 全局App
var (
	// DoApp *App
	StaticDir map[string]string

)
// 初始化
func init() {
	// DoApp = NewApp()
	StaticDir = make(map[string]string)
	StaticDir["/favicon.ico"] = "/public/favicon.ico"
}
// 定义结构
type App struct {
	Handler		*RouteMap
	Server 		*http.Server
}
// 生成App
func NewApp(cfg *Config) *App{
	return &App{
		Handler:nil,
		Server:nil,
	}
}
//注册Server
func (a *App)RegisterServer(server *http.Server) {
	a.Server = server
}
//注册Handler
func (a *App)RegisterHandler(handler *RouteMap) {
	a.Handler = handler
}

func SetStaticDir(url string,path string) {
	StaticDir[url] = path
}

func (a *App)Run(){
	// Log.Info("%v",a.Handler)
	a.Server.Handler = a.Handler
	err := a.Server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}