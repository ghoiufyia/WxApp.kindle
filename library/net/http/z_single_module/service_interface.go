package modules

import (
	"github.com/gorilla/mux"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
	// Exported methods
	GetRoutes() []Route
	RegisterRoutes(router *mux.Router, prefix string)
	Close()
}