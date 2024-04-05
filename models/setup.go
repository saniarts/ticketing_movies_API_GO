package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=12345 dbname=ticketing_test port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&City{})
	database.AutoMigrate(&Cinema{})
	database.AutoMigrate(&CinemaScreen{})
	database.AutoMigrate(&Movie{})

	DB = database
}
