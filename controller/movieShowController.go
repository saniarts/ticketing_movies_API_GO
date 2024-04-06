package controller

import (
	"net/http"
	"ticketing_movies_API_GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieShowController struct{}

func (ms *MovieShowController) Index(c *gin.Context) {
	var movieShows []models.MovieShow

	if err := models.DB.Preload("Movie").Preload("CinemaScreen").Find(&movieShows).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch movie shows"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Movie Shows": movieShows})

}

func (ms *MovieShowController) Show(c *gin.Context) {
	var movieShow models.MovieShow

	id := c.Param("id")

	if err := models.DB.Preload("Movie").Preload("CinemaScreen").First(&movieShow, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Movie Show": movieShow})
}

func (ms *MovieShowController) Create(c *gin.Context) {

	var movieShow models.MovieShow

	if err := c.ShouldBindJSON(&movieShow); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var cinemaScreen models.CinemaScreen
	if err := models.DB.Preload("Cinema").Preload("Cinema.City").First(&cinemaScreen, movieShow.CinemaScreenID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cinema not found"})
		return
	}
	movieShow.CinemaScreen = cinemaScreen

	var movie models.Movie
	if err := models.DB.First(&movie, movieShow.MovieID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Movie not found"})
		return
	}
	movieShow.Movie = movie

	if err := models.DB.Create(&movieShow).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Movie Show": movieShow})

}

func (ms *MovieShowController) Update(c *gin.Context) {
	var movieShow models.MovieShow

	id := c.Param("id")

	if err := c.ShouldBindJSON(&movieShow); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var cinemaScreen models.CinemaScreen
	if err := models.DB.Preload("Cinema").Preload("Cinema.City").First(&cinemaScreen, movieShow.CinemaScreenID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cinema not found"})
		return
	}

	var movie models.Movie
	if err := models.DB.First(&movie, movieShow.MovieID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Movie not found"})
		return
	}

	if models.DB.Model(&movieShow).Where("id = ?", id).Updates(&movieShow).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func (ms *MovieShowController) Delete(c *gin.Context) {

	var movieShow models.MovieShow

	var input struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.First(&movieShow, input.ID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&movieShow).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
