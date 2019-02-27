package app

import (
	// "github.com/gorilla/mux"
	// "net/http"
	"fmt"
)

type App struct {
	// Server 		*http.Server
	Routes		*RegisterRoutes
}

func NewApp() *App{
	routes := NewRegisterRoutes()
	return &App{Routes:routes}
}

func (a *App)Run(){
	fmt.Printf("%v",a.Routes)
}