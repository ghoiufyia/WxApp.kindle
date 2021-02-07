package controllers

import (
	"fmt"
	// "github.com/ghoiufyia/WxApp.kindle/web/models"
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	// "log"
	// email_pb "github.com/ghoiufyia/WxApp.kindle/email_service/proto/email"
	// "google.golang.org/grpc"
	// "context"
	"github.com/ghoiufyia/kindle/app/web/kindle/internal/resource"
	"github.com/ghoiufyia/kindle/library/net/http/klove"

)

type IndexController struct {
	klove.Controller
	name	string
	age		int32
}

func (i *IndexController)Index ()  {
	res := resource.Get()

	useremail,err := res.First()
	fmt.Printf("%+v\n",useremail)
	if err != nil {
		fmt.Errorf("first error is %+v.",err)
	}
	
	i.Ctx.RenderTemplate("index","base","sssss")
}

func (i *IndexController) Json ()  {
	data := make(map[string]string, 0)
	data["code"] = "200"
	data["msg"] = "ok"
	i.Ctx.RenderJson(data)
}

// func (i *IndexController)Session ()  {
// 	sessionid := i.Ctx.Session.GetSid()
// 	i.Ctx.Session.Set("login_token","1")

// 	i.RenderString(sessionid)
// }


// func (i *IndexController)Json ()  {
// 	db := dogo.GetDB()
// 	var user_email models.UserEmail
// 	db.First(&user_email)

// 	i.SetData("code","1")
// 	i.SetData("msg","ok")
// 	// i.SetData("data",make(map[string]string, 0))
// 	i.SetData("data",user_email)
// 	i.RenderJson()
// }

// func (i *IndexController)Session ()  {
// 	sessionid := i.Ctx.Session.GetSid()
// 	i.Ctx.Session.Set("login_token","1")

// 	i.RenderString(sessionid)
// }



// func (i *IndexController)List ()  {

// 	// i.Render()
// 	i.RenderTemplate("list.html")
// }

// const (
// 	ADDRESS           = "127.0.0.1:8000"
// )

// func (i *IndexController)Rpc ()  {
// 	// 连接远端服务
// 	conn,err := grpc.Dial(ADDRESS,grpc.WithInsecure())
// 	if err != nil {
// 		dogo.Info("connect error %v",err)
// 	}
// 	defer conn.Close()
// 	// 定义客户端
// 	client := email_pb.NewEmailServiceClient(conn)

// 	var resuest = email_pb.CreateEmailRequest{UserId:2,Email:"ddd@11.com"}

// 	// 调用 RPC
// 	resp, err := client.CreateEmail(context.Background(), &resuest)
// 	if err != nil {
// 		dogo.Info("create email error: %v", err)
// 	}

// 	dogo.Info("created: %+v", resp)
// 	conn.Close()

// 	// i.Render()


// }