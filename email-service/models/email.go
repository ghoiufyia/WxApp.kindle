package models

import (
	// "database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	gorm.Model
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type TimestampModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserEmail struct {
	BaseModel
	UserId		string `gorm:"unique;not null;default ''"`	
	Email		string `gorm:"type:varchar(100)"`
}

type Email struct {
	BaseModel
	UserId		string `gorm:"unique;not null;default ''"`	
	ToEmail		string `gorm:"type:varchar(100)"`
	FilePath	string `gorm:"type:varchar(255)"`
}
// create table `user_email` (
// 	`id` varchar(255) not null primary key,
// 	`user_id` varchar(255) not null default '',
// 	`email` varchar(100) not null default '',
// 	`created_at` int(11) not null default 0,
// 	`updated_at` int(11) not null default 0,
// 	`deleted_at` int(11) not null default 0,
// 	unique key (`user_id`),
// 	unique key (`email`)
// 	)engine=innodb character set utf8mb4; 

