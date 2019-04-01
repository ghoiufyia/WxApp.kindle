package databse

import (
	"fmt"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

type DatabaseConfig struct {
	Type         string `default:"mysql"`
	Host         string `default:"localhost"`
	Port         int    `default:"3306"`
	User         string `default:"gravitee"`
	Password     string `default:"gravitee"`
	DatabaseName string `default:"gravitee"`
	MaxIdleConns int    `default:"5"`
	MaxOpenConns int    `default:"5"`
}

var (
	DoDB *gorm.DB
)

func InitDatabse(cfg DatabaseConfig) error{
	var err error
	DoDB,err = NewDB(cfg)
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	if nil != DoDB {
		return DoDB
	}
	cfg := NewConfig("")
	var err error
	DoDB,err = NewDB(cfg.Database)
	if err != nil {
		Info("DoDB error:%+v",DoDB)
	}
	return DoDB
}

func NewDB(cfg DatabaseConfig) (*gorm.DB,error) {
	if cfg.Type == "mysql" {
		args := fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DatabaseName,
		)
		db,err := gorm.Open(cfg.Type,args)
		if err != nil {
			return db,err
		}
		// Max idle connections
		db.DB().SetMaxIdleConns(cfg.MaxIdleConns)

		// Max open connections
		db.DB().SetMaxOpenConns(cfg.MaxOpenConns)

		// Database logging
		db.LogMode(false)

		return db, nil
	}
	return nil, fmt.Errorf("Database type %s not supported", cfg.Type)
}