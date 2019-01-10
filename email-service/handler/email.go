package handler

import (
	email "WxApp.kindle/email-service/proto"
	"context"

)

type service struct {

}

func (s *service) Create(ctx context.Context,req *email.CreateRequest) (*email.CreateResponse,error) {


}