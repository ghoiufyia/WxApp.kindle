package controllers

import (
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
	"fmt"
	// "github.com/ghoiufyia/WxApp.kindle/web/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type IndexController struct {
	dogo.Controller
	Name	string
	age		int32
}

func (i *IndexController)Json ()  {
	data := map[string]interface{}{
			"msg": "ok",
		}
		args := fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DatabaseName,
		)
	db, err := sql.Open("mysql", args)
	rows, err := db.Query("SELECT * FROM userinfo")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
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

