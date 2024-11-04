package models

import (
	"log"
	"time"
)

type UserDB struct {
	ID              int       `gorm:"primaryKey"`
	Name            string    `gorm:"type:varchar(255)"`
	Email           string    `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHash    string    `gorm:"type:varchar(255)"`
	Occupation      string    `gorm:"type:varchar(255)"`
	Avatar_filename string    `gorm:"type:varchar(255)"`
	Role            string    `gorm:"type:varchar(255)"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

func (UserDB) tableName() string {
	log.Println("table name ke eksekusi")
	return "users"
}
