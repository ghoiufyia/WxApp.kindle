package models

import (
	// "github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type UserEmail struct {
	BaseModel
	UserId		uint `gorm:"not null;default ''"`	
	Email		string `gorm:"type:varchar(100)"`
}

type Email struct {
	BaseModel
	UserId		uint `gorm:"not null;default ''"`	
	ToEmail		string `gorm:"type:varchar(100)"`
	FilePath	string `gorm:"type:varchar(255)"`
}

