package resource

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/ghoiufyia/kindle/app/web/kindle/conf"
	"github.com/ghoiufyia/kindle/library/database/mysql"

)


var (
	res *Resource
)

type Resource struct {
	config            *conf.Config
	// authMC       *memcache.Pool
	// authMCExpire int32
	originDB     *gorm.DB
}

func New(c *conf.Config) (err error) {
	db,err := mysql.NewMysql(c.OrmConf)
	if err != nil {
		return
	}
	res = &Resource{
		config : c,
		originDB :	db,
	}
	return
}

func Get() (r *Resource){
	return res
}