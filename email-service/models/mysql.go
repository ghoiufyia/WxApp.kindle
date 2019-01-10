package models

import (
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // 导入数据库驱动
	email "WxApp.kindle/email-service/proto"
)

func init()  {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterModel(new(email.Email))

}
