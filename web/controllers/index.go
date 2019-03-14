package controllers

import (
	// "net/http"
	// "encoding/json"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"fmt"
	// "io"
	"net/http"
	"encoding/json"
	email_pb "github.com/ghoiufyia/WxApp.kindle/email-service/proto/email"
	"google.golang.org/grpc"
	"log"
	"context"
)

type IndexController struct {
	dogo.Controller
	Name	string
	age		int32
}

func (i *IndexController)Index ()  {
	fmt.Printf("adsd===============================")
	// fmt.Println(i.Ctx)
	fmt.Printf("%v",i.Ctx.Request)
	// io.WriteString(i.Ctx.ResponseWriter, "Hello from a HandleFunc #2!\n")
	
	fmt.Printf("adsd===============================")
	fmt.Printf("adsd===============================")

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

	var resuest = email_pb.CreateEmailRequest{UserId:"fghfyffffffffjy",Email:"ddd@11.com"}

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









// func (i *IndexController)Index (w http.ResponseWriter, r *http.Request)  {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"msg": "ok",
// 	})
// }

// func (ih *IndexHandler) SetAge (age int32) (error) {
// 	ih.age = age
// 	return nil
// }

// func (ih *IndexHandler) GetAge () (int32,error) {
// 	return ih.age,nil
// }
