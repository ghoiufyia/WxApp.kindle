package cmd

import (
	"WxApp.kindle/email-service/models"
)
//弃用
func Migrate(configFile string) error {
	cfg := config.NewConfig(configFile)
	db,err := databse.NewDatabase(cfg)
	if err != nil {
		return nil
	}
	defer db.Close()

}