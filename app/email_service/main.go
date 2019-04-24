package main

import (
	"google.golang.org/grpc"
	"net"
	"github.com/ghoiufyia/WxApp.kindle/email_service/config"
	"github.com/ghoiufyia/WxApp.kindle/email_service/database"
	"log"
	"github.com/ghoiufyia/WxApp.kindle/email_service/services"
	email_pb "github.com/ghoiufyia/WxApp.kindle/email_service/proto/email"
	"strconv"
)



func main()  {
	configFile := ""
	cfg := config.NewConfig(configFile)
	port := cfg.ServerPort

	listener,err := net.Listen("tcp",":"+strconv.Itoa(cfg.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}
	log.Printf("listen on: %d\n", port)
	server := grpc.NewServer()
	db,err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to open database: %v",err)
	}
	log.Printf("listen on: %s\n", "database is ok")
	emailServer := &services.EmailServer{db}
	email_pb.RegisterEmailServiceServer(server,emailServer)
	if err := server.Serve(listener);err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}