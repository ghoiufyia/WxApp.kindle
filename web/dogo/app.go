package dogo

import (
	// "github.com/gorilla/mux"
	"net/http"
	"fmt"
)

// 全局App
// var (
// 	DoApp *App
// )
// 初始化
func init() {
	// DoApp = NewApp()
}
// 定义结构
type App struct {
	Handler		*RouteMap
	Server 		*http.Server
}
// 生成App
func NewApp(cfg *Config) *App{
	//加载路由
	rm := NewRouteMap()
	return &App{
		Handler:rm,
		Server:&http.Server{
			Addr:":8080",
		},
	}
}

func (a *App)Run(){
	fmt.Printf("%v",a.Handler)
	// a.Server.Handler = a.Handler
	a.Server.Handler = a.Handler
	err := a.Server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}

func Router(prefix string,c ControllerInterface) {
	DoApp.Handler.RegisterRouteGroup(prefix,c)
}