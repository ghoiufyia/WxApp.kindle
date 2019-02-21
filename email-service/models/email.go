package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/RichardKnop/uuid"
)

type BaseModel struct {
	gorm.model
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type TimestampModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserEmail struct {
	BaseModel
	UserId		string `gorm:"unique;not null;default ''"`	
	ToEmail		string `gorm:"type:varchar(100)"`
}

type Email struct {
	BaseModel
	UserId		string `gorm:"unique;not null;default ''"`	
	ToEmail		string `gorm:"type:varchar(100)"`
	FilePath	string `gorm:"type:varchar(255)"`
}


