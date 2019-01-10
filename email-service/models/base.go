package models

import (
	"github.com/astaxie/beego/orm"
	email "WxApp.kindle/email-service/proto"
)

func init()  {
	orm.RegisterModel(new(email.Email))
}
