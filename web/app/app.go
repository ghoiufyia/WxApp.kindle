package app

import (
	// "github.com/gorilla/mux"
	"net/http"
	"fmt"
)

type App struct {
	Handlers		*RegisterRoutes
	Server 		*http.Server
}

func NewApp() *App{
	routes := NewRegisterRoutes()
	return &App{Handlers:routes,Server:&http.Server{}}
}

func (a *App)Run(){
	fmt.Printf("%v",a.Handlers)
	a.Server.Handler = a.Handlers

}