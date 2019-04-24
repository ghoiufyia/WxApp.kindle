package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
)

type Email struct {
	dogo.BaseModel
	UserId		uint `gorm:"not null;default ''"`	
	ToEmail		string `gorm:"type:varchar(100)"`
	FilePath	string `gorm:"type:varchar(255)"`
}

