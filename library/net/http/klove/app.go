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
	StaticDir["/public"] = "./public"
}
// 定义结构
type App struct {
	Server 		*http.Server
}
// 生成App
func NewApp(cfg *Config) *App{
	return &App{
		Server:nil,
	}
}
//注册Server
func (a *App)RegisterServer(server *http.Server) {
	a.Server = server
}


func SetStaticDir(url string,path string) {
	StaticDir[url] = path
}

func (a *App)Run(){
	// Log.Info("%v",a.Handler)
	// a.Server.Handler = a.Handler
	err := a.Server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}