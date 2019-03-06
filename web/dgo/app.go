package dgo

import (
	// "github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var (
	DApp *App
)

func init() {
	DApp = NewApp()
}

type App struct {
	Handlers	*RegisterRoutes
	Server 		*http.Server
}

func NewApp() *App{
	routes := NewRegisterRoutes()
	return &App{Handlers:routes,Server:&http.Server{}}
}

func (a *App)Run(){
	fmt.Printf("%v",a.Handlers)
	// a.Server.Handler = a.Handlers

}

// app.Server.Handler = app.Handlers
// 	for i := len(mws) - 1; i >= 0; i-- {
// 		if mws[i] == nil {
// 			continue
// 		}
// 		app.Server.Handler = mws[i](app.Server.Handler)
// 	}