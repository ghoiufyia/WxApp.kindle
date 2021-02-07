package modules

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

// RegisterRoutes registers route handlers for the health service
func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	addRoutes(s.GetRoutes(), subRouter)
}

// GetRoutes returns []routes.Route slice for the health service
func (s *Service) GetRoutes() []Route {
	return []Route{
		{
			Name:        "health_check",
			Method:      "GET",
			Pattern:     "/index",
			HandlerFunc: s.index,
		},
	}
}

func addRoutes(routes []Route,router *mux.Router) {
	if len(routes) > 0 {
		for _,route := range routes {
			router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}
}