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
	// "github.com/ghoiufyia/kindle/app/web/kindle/internal/resource"
	"github.com/ghoiufyia/kindle/library/net/http/klove"

)

type KindleController struct {
	klove.Controller
	name	string
	age		int32
}

func (i *KindleController) KBindForm ()  {
	// res := resource.Get()

	// useremail,err := res.First()
	// fmt.Printf("%+v\n",useremail)
	// if err != nil {
	// 	fmt.Errorf("first error is %+v.",err)
	// }
	// fmt.Printf("%+v\n",i.Ctx.Request.Method)
	// fmt.Printf("%+v\n",i.Ctx.Request)
	// i.Ctx.RenderJson(i.Ctx.Request)

	i.Ctx.RenderTemplate("kbind_form","base","sssss")
}

func (i *KindleController) KBind ()  {
	i.Ctx.RenderJson(i.Ctx.Request)


	// res := resource.Get()

	// useremail,err := res.First()
	// fmt.Printf("%+v\n",useremail)
	// if err != nil {
	// 	fmt.Errorf("first error is %+v.",err)
	// }
	
	// i.Ctx.RenderTemplate("index","base","sssss")
}