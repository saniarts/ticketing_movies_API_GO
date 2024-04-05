package models

import (
	"time"

	"gorm.io/gorm"
)

type MovieShow struct {
	gorm.Model
	StartTime      time.Time    `gorm:"not null" json:"StartTime"`
	EndTime        time.Time    `gorm:"not null" json:"EndTime"`
	CinemaScreenID uint         `gorm:"not null" json:"CinemaScreenID"`
	CinemaScreen   CinemaScreen `gorm:"foreignKey:CinemaScreenID" json:"CinemaScreen"`
	MovieID        uint         `gorm:"not null" json:"MovieID"`
	Movie          Movie        `gorm:"foreignKey:MovieID" json:"Movie"`
}
