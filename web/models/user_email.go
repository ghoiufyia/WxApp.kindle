package models

import (
	// "github.com/jinzhu/gorm"
	"github.com/ghoiufyia/WxApp.kindle/web/dogo"
)

type UserEmail struct {
	dogo.BaseModel
	UserId		uint32 `gorm:"not null;default ''"`	
	Email		string `gorm:"type:varchar(100)"`
}
