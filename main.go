package main

import (
	controller "ticketing_movies_API_GO/controller"
	"ticketing_movies_API_GO/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	cityController := &controller.CityController{}
	r.GET("/api/cities", cityController.Index)
	r.GET("/api/city/:id", cityController.Show)
	r.POST("/api/city", cityController.Create)
	r.PUT("/api/city/:id", cityController.Update)
	r.DELETE("/api/city", cityController.Delete)

	cinemaController := &controller.CinemaController{}
	r.GET("/api/cinemas", cinemaController.Index)
	r.GET("/api/cinema/:id", cinemaController.Show)
	r.POST("/api/cinema", cinemaController.Create)
	r.PUT("/api/cinema/:id", cinemaController.Update)
	r.DELETE("/api/cinema", cinemaController.Delete)

	cinemaScreenController := &controller.CinemaScreenController{}
	r.GET("/api/cinema_screens", cinemaScreenController.Index)
	r.GET("/api/cinema_screen/:id", cinemaScreenController.Show)
	r.POST("/api/cinema_screen", cinemaScreenController.Create)
	r.PUT("/api/cinema_screen/:id", cinemaScreenController.Update)
	r.DELETE("/api/cinema_screen", cinemaScreenController.Delete)

	movieController := &controller.MovieController{}
	r.GET("/api/movies", movieController.Index)
	r.GET("/api/movie/:id", movieController.Show)
	r.POST("/api/movie", movieController.Create)
	r.PUT("/api/movie/:id", movieController.Update)
	r.DELETE("/api/movie", movieController.Delete)

	movieShowController := &controller.MovieShowController{}
	r.GET("/api/movie_shows", movieShowController.Index)
	r.GET("/api/movie_show/:id", movieShowController.Show)
	r.POST("/api/movie_show", movieShowController.Create)
	r.PUT("/api/movie_show/:id", movieShowController.Update)
	r.DELETE("/api/movie_show", movieShowController.Delete)

	r.Run()
}
