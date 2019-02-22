package main

import (
	"github.com/gorilla/mux"
	"net/http"

)

type App struct {
	Server 		*http.Server
	Route		*mux.Router
	
}

func app()  {
	
}