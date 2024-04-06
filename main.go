package main

import (
	controller "ticketing_movies_API_GO/controller"
	"ticketing_movies_API_GO/initializers"
	"ticketing_movies_API_GO/middleware"
	"ticketing_movies_API_GO/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	cityController := &controller.CityController{}
	r.GET("/api/cities", middleware.RequireAuth, cityController.Index)
	r.GET("/api/city/:id", middleware.RequireAuth, cityController.Show)
	r.POST("/api/city", middleware.RequireAuth, cityController.Create)
	r.PUT("/api/city/:id", middleware.RequireAuth, cityController.Update)
	r.DELETE("/api/city", middleware.RequireAuth, cityController.Delete)

	cinemaController := &controller.CinemaController{}
	r.GET("/api/cinemas", middleware.RequireAuth, cinemaController.Index)
	r.GET("/api/cinema/:id", middleware.RequireAuth, cinemaController.Show)
	r.POST("/api/cinema", middleware.RequireAuth, cinemaController.Create)
	r.PUT("/api/cinema/:id", middleware.RequireAuth, cinemaController.Update)
	r.DELETE("/api/cinema", middleware.RequireAuth, cinemaController.Delete)

	cinemaScreenController := &controller.CinemaScreenController{}
	r.GET("/api/cinema_screens", middleware.RequireAuth, cinemaScreenController.Index)
	r.GET("/api/cinema_screen/:id", middleware.RequireAuth, cinemaScreenController.Show)
	r.POST("/api/cinema_screen", middleware.RequireAuth, cinemaScreenController.Create)
	r.PUT("/api/cinema_screen/:id", middleware.RequireAuth, cinemaScreenController.Update)
	r.DELETE("/api/cinema_screen", middleware.RequireAuth, cinemaScreenController.Delete)

	movieController := &controller.MovieController{}
	r.GET("/api/movies", middleware.RequireAuth, movieController.Index)
	r.GET("/api/movie/:id", middleware.RequireAuth, movieController.Show)
	r.POST("/api/movie", middleware.RequireAuth, movieController.Create)
	r.PUT("/api/movie/:id", middleware.RequireAuth, movieController.Update)
	r.DELETE("/api/movie", middleware.RequireAuth, movieController.Delete)

	movieShowController := &controller.MovieShowController{}
	r.GET("/api/movie_shows", middleware.RequireAuth, movieShowController.Index)
	r.GET("/api/movie_show/:id", middleware.RequireAuth, movieShowController.Show)
	r.POST("/api/movie_show", middleware.RequireAuth, movieShowController.Create)
	r.PUT("/api/movie_show/:id", middleware.RequireAuth, movieShowController.Update)
	r.DELETE("/api/movie_show", middleware.RequireAuth, movieShowController.Delete)

	authController := &controller.AuthController{}
	r.POST("/api/register", authController.Register)
	r.POST("/api/login", authController.Login)

	r.Run()
}
