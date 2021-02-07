package dogo

import (
	"net/http"
	"strconv"
)
//生成新Server
func NewServer(cfg ServerConfig) (*http.Server,error) {
	serverPort := cfg.ServerPort
	var addr string = ":"+strconv.Itoa(serverPort)
	return &http.Server{
		Addr:addr,
	},nil
}

//注册Handler
func SetServerHandler(server *http.Server,handler *RouteMap) {
	server.Handler = handler
}