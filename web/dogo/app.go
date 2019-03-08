package dogo

import (
	// "github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var (
	DoApp *App
)

func init() {
	DoApp = NewApp()
}

type App struct {
	Handler		*RouteMap
	Server 		*http.Server
}

func NewApp() *App{
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





// app.Server.Handler = app.Handlers
// 	for i := len(mws) - 1; i >= 0; i-- {
// 		if mws[i] == nil {
// 			continue
// 		}
// 		app.Server.Handler = mws[i](app.Server.Handler)
// 	}