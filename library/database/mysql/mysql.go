package mysql

import (
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	Type         string `default:"mysql"`
	Host         string `default:"localhost"`
	Port         int    `default:"3306"`
	User         string `default:"gravitee"`
	Password     string `default:"gravitee"`
	DatabaseName string `default:"gravitee"`
	MaxIdleConns int    `default:"5"`
	MaxOpenConns int    `default:"5"`
}

func NewMysql(cfg *Config) (*gorm.DB,error) {
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