package models

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);not null" json:"Name"`
	Zipcode string `gorm:"type:varchar(255);not null" json:"Zipcode"`
}
