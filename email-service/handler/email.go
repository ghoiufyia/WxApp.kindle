package handler

import (
	email "WxApp.kindle/email-service/proto"
	"context"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type service struct {

}

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}


func (s *service) Create(ctx context.Context,req *email.CreateRequest) (*email.CreateResponse,error) {


}

func (s *service) Select(ctx context.Context, req *email.SelectRequest, out *SelectResponse) error {
	return h.EmailServiceHandler.Select(ctx, in, out)
}

func (s *service) Delete(ctx context.Context, req *email.DeleteRequest, out *DeleteResponse) error {
	return h.EmailServiceHandler.Delete(ctx, in, out)
}