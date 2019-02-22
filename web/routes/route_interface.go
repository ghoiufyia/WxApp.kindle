package routes

import (
	"github.com/gorilla/mux"
)

type RoutesInterface interface {
	GetRoutes() []Route
	RegisterRoutes(router *mux.Router)
	Close()
}