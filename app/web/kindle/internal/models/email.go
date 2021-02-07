package models

import (
)

type Email struct {
	UserId		uint `gorm:"not null;default ''"`	
	ToEmail		string `gorm:"type:varchar(100)"`
	FilePath	string `gorm:"type:varchar(255)"`
}

