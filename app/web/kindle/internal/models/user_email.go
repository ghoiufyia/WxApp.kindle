package models

import (
)

type UserEmail struct {
	UserId		uint32 `gorm:"not null;default ''"`	
	Email		string `gorm:"type:varchar(100)"`
}
