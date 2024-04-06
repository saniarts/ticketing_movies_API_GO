package controller

import (
	"net/http"
	"ticketing_movies_API_GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CinemaScreenController struct{}

func (cs *CinemaScreenController) Index(c *gin.Context) {
	var cinemascreens []models.CinemaScreen

	if err := models.DB.Preload("Cinema").Preload("Cinema.City").Find(&cinemascreens).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch cinemascreens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Cinema Screens": cinemascreens})

}

func (cs *CinemaScreenController) Show(c *gin.Context) {
	var cinemascreen models.CinemaScreen

	id := c.Param("id")

	if err := models.DB.Preload("Cinema").Preload("Cinema.City").First(&cinemascreen, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Cinema Screen": cinemascreen})
}

func (cs *CinemaScreenController) Create(c *gin.Context) {

	var cinemascreen models.CinemaScreen

	if err := c.ShouldBindJSON(&cinemascreen); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var cinema models.Cinema
	if err := models.DB.Preload("City").First(&cinema, cinemascreen.CinemaID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cinema not found"})
		return
	}

	cinemascreen.Cinema = cinema
	if err := models.DB.Create(&cinemascreen).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Cinema Screen": cinemascreen})

}

func (cs *CinemaScreenController) Update(c *gin.Context) {
	var cinemascreen models.CinemaScreen

	id := c.Param("id")

	if err := c.ShouldBindJSON(&cinemascreen); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var cinema models.Cinema
	if err := models.DB.Preload("City").First(&cinema, cinemascreen.CinemaID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cinema not found"})
		return
	}

	if models.DB.Model(&cinemascreen).Where("id = ?", id).Updates(&cinemascreen).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func (cs *CinemaScreenController) Delete(c *gin.Context) {

	var cinemascreen models.CinemaScreen

	var input struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.First(&cinemascreen, input.ID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&cinemascreen).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
