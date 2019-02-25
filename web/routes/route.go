package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
}

type Routes struct {
	Prefix		string
	MyRoutes 	[]Route
}


func (r *Routes)GetRoutes() []Route {
	return r.MyRoutes
}

func (r *Routes)RegisterRoutes(router *mux.Router) {
	subRouter := router.PathPrefix(r.Prefix).Subrouter()
	addRoute(r.GetRoutes(),router)
}

func addRoute(routes []Route,router *mux.Router) {
	if len(routes) > 0 {
		for _,route := range routes {
			router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}
}