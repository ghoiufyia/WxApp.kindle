package dgo

import (
	// "net/http"
	// "github.com/gorilla/mux"
	"reflect"
)

type Route struct {
	Name		string
	Method		string
	Pattern		string
	Controller	reflect.Type
	Action		string
}

type RouteGroup struct {
	Prefix		string
	MyRoutes 	[]*Route
}

type RegisterRoutes struct {
	Routes map[string]RouteGroup
}

func NewRegisterRoutes() *RegisterRoutes {
	return &RegisterRoutes{Routes:make(map[string]RouteGroup)}
}

// func (r *Routes)GetRoutes() []Route {
// 	return r.MyRoutes
// }

// func (r *Routes)RegisterRoutes(router *mux.Router) {
// 	subRouter := router.PathPrefix(r.Prefix).Subrouter()
// 	addRoute(r.GetRoutes(),router)
// }

// func addRoute(routes []Route,router *mux.Router) {
// 	if len(routes) > 0 {
// 		for _,route := range routes {
// 			router.Methods(route.Method).
// 				Path(route.Pattern).
// 				Name(route.Name).
// 				Handler(route.HandlerFunc)
// 		}
// 	}
// }