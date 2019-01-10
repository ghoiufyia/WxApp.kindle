package handler

import (
	email "WxApp.kindle/email-service/proto"
	"net"
	"context"
	"log"
	"google.golang.org/grpc"
)

type service struct {

}

func (s *service) Create(ctx context.Context,req *email.CreateRepuest) (*email.CreateResponse,error) {
	

}