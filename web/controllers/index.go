package controllers

import (
	// "net/http"
	// "encoding/json"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"fmt"
	// "io"
	// "net/http"
	// "encoding/json"
	// email_pb "github.com/ghoiufyia/WxApp.kindle/email-service/proto/email"
	// "google.golang.org/grpc"
	// "log"
	// "context"
)

type IndexController struct {
	dogo.Controller
	// DB *gorm.DB
	Name	string
	age		int32
}

func (i *IndexController)Json ()  {
	fmt.Printf("json========json=======================")
	data := map[string]interface{}{
			"msg": "ok",
		}
	i.Data = data
	i.SetData("code","1")
	i.SetData("msg","ok")
	i.SetData("data",make(map[string]string, 0))
	i.RenderJson()
}


func (i *IndexController)Index ()  {
	fmt.Printf("adsd===============================")
	// fmt.Println(i.Ctx)
	// fmt.Printf("%+v",i.Ctx.Request)

	fmt.Printf("adsd========ffffff=======================")

	i.Render()

	// // 连接远端服务
	// conn,err := grpc.Dial(ADDRESS,grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("connect error %v",err)
	// }
	// defer conn.Close()
	// // 定义客户端
	// client := email_pb.NewEmailServiceClient(conn)

	// var resuest = email_pb.CreateEmailRequest{UserId:"fghfyffffffffjy",Email:"ddd@11.com"}

	// // 调用 RPC
	// resp, err := client.CreateEmail(context.Background(), &resuest)
	// if err != nil {
	// 	log.Printf("create email error: %v", err)
	// }

	// log.Printf("created: %t", resp.Msg)
	// conn.Close()

}

