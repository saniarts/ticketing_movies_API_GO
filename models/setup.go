package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&City{})
	database.AutoMigrate(&Cinema{})
	database.AutoMigrate(&CinemaScreen{})
	database.AutoMigrate(&Movie{})
	database.AutoMigrate(&MovieShow{})
	database.AutoMigrate(&User{})

	DB = database
}
