package models

import (
	"gorm.io/gorm"
)

type Cinema struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255);not null" json:"Name"`
	CityID uint   `gorm:"not null" json:"CityID"`
	City   City   `gorm:"foreignKey:CityID" json:"City"`
}
