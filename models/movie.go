package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null" json:"Title"`
	Description string `gorm:"type:varchar(255);not null" json:"Description"`
	Duration    int    `gorm:"not null" json:"Duration"`
}
