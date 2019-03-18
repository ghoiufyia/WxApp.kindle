package dogo

import (
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	// "errors"
)

var (
	DoDB *gorm.DB
)


func GetDB() *gorm.DB {
	if nil != DoDB {
		return DoDB
	}
	cfg := NewConfig("")
	var err error
	DoDB,err = NewDB(cfg.Database)
	if err != nil {
		Log.Info("DoDB error:%+v",DoDB)
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
		db.LogMode(true)

		return db, nil
	}
	return nil, fmt.Errorf("Database type %s not supported", cfg.Type)
}