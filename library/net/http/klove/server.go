package klove

import (
	"net/http"
	"strconv"
	"fmt"
)
//server config
type ServerConfig struct {
	ServerAddr		string
	ServerPort		int  `default:"8080"`
}
//
type HandlerFunc func(*Context)
//引擎，控制应用流程
type Engine struct {
	RouteGroup
	middleware	[]HandlerFunc
	conf		*ServerConfig
	server 		*http.Server
	handler 	http.Handler
}

//new Server
func NewWebEngine(conf *ServerConfig) (*Engine) {
	engine := &Engine{
		RouteGroup:	RouteGroup{
			prefix:		"",
			routes:		make(map[string]route,0),
		},
		middleware:		make([]HandlerFunc,0),
		conf:			conf,
	}
	return engine
}

//set engine Handler,use for server handler
func (e *Engine)SetHandler(handler http.Handler){
	e.handler = handler
}

//init http server，store in engine，open goroutine and run
func (e *Engine)Start() error{
	conf := e.conf
	serverAddr := conf.ServerAddr
	serverPort := conf.ServerPort
	handler := e.handler
	var addr string = serverAddr + ":" + strconv.Itoa(serverPort)
	server := &http.Server{
		Addr:addr,
		Handler:handler,
	}
	fmt.Printf("server is serve on %s...\n",addr)
	//开线程轮询server
	go func(){
		e.server = server
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	return nil
}