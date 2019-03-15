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