package models

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)

    orm.RegisterDataBase("poetnote", "mysql", "root:Sn93007997@tcp(101.132.37.148:3306)/poetnote?charset=utf8")
}
