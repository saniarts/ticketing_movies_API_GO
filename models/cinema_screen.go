package models

import (
	"gorm.io/gorm"
)

type CinemaScreen struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"Name"`
	CinemaID uint   `gorm:"not null" json:"CinemaID"`
	Cinema   Cinema `gorm:"foreignKey:CinemaID" json:"Cinema"`
}
