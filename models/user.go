package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"Name"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"Email"`
	Password string `gorm:"type:varchar(255);not null" json:"Password"`
}
