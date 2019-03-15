package dogo

import (
	// "net/http"
	// "github.com/gorilla/mux"
	"reflect"
	"net/http"
	// "fmt"
	// "context"
)
//单个路由
type route struct {
	Name		string
	Method		string
	Pattern		string
	Controller	reflect.Type
	Action		string
}
//一个路由组
type RouteGroup struct {
	Prefix		string
	Controller 	reflect.Type
}
//多个路由组
type RouteMap struct {
	Routes []*RouteGroup
}
//新生成路由组
func NewRouteMap() *RouteMap {
	return &RouteMap{
		Routes:make([]*RouteGroup, 0),
	}
}
//添加路由
func (rm *RouteMap)RegisterRouteGroup(prefix string,c ControllerInterface) {
	controller := reflect.Indirect(reflect.ValueOf(c)).Type()
	rg := &RouteGroup{}
	rg.Prefix = prefix
	rg.Controller = controller
	rm.Routes = append(rm.Routes,rg)
}
func (rm *RouteMap)Router(prefix string,c ControllerInterface) {
	rm.RegisterRouteGroup(prefix,c)
}


//实现http的server handler
func (rm *RouteMap)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vc := reflect.New(rm.Routes[0].Controller)
	init := vc.MethodByName("Init")
	in := make([]reflect.Value, 1)
	ctx := &Context{
		ResponseWriter:w,
		Request:r,
	}
	in[0] = reflect.ValueOf(ctx)
	init.Call(in)
	in = make([]reflect.Value, 0)
	index := vc.MethodByName("Index")
	index.Call(in)

	// render := vc.MethodByName("Render")
	// render.Call(in)

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