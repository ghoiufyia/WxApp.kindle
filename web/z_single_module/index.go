package modules

import (
	"net/http"
	"encoding/json"
	email_pb "github.com/ghoiufyia/WxApp.kindle/email-service/proto/email"
	"google.golang.org/grpc"
	"log"
	"context"
)

const (
	ADDRESS           = "127.0.0.1:8000"
	DEFAULT_INFO_FILE = "order.json"
)

func (s *Service) index (w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)

	// 连接远端服务
	conn,err := grpc.Dial(ADDRESS,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error %v",err)
	}
	defer conn.Close()
	// 定义客户端
	client := email_pb.NewEmailServiceClient(conn)

	var resuest = email_pb.CreateEmailRequest{UserId:2,Email:"ddd@11.com"}

	// 调用 RPC
	resp, err := client.CreateEmail(context.Background(), &resuest)
	if err != nil {
		log.Printf("create email error: %v", err)
	}

	log.Printf("created: %t", resp.Msg)
	conn.Close()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg": "ok",
	})
}

func (s *Service) SetAge (age int32) (error) {
	s.age = age
	return nil
}

func (s *Service) GetAge () (int32,error) {
	return s.age,nil
}