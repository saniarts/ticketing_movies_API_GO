package controller

import (
	"fmt"
	"net/http"
	"ticketing_movies_API_GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CinemaController struct{}

func (cn *CinemaController) Index(c *gin.Context) {
	var cinemas []models.Cinema

	if err := models.DB.Preload("City").Find(&cinemas).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Cinemas": cinemas})

}

func (cn *CinemaController) Show(c *gin.Context) {
	var cinema models.Cinema

	id := c.Param("id")

	if err := models.DB.Preload("City").First(&cinema, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Cinema": cinema})
}

func (cn *CinemaController) Create(c *gin.Context) {

	var cinema models.Cinema

	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var city models.City
	if err := models.DB.First(&city, cinema.CityID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "City not found"})
		return
	}

	cinema.City = city
	if err := models.DB.Create(&cinema).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Cinema": cinema})

}

func (cn *CinemaController) Update(c *gin.Context) {
	var cinema models.Cinema

	id := c.Param("id")

	fmt.Println("Before binding JSON:", &cinema)

	if err := c.ShouldBindJSON(&cinema); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var city models.City
	if err := models.DB.First(&city, cinema.CityID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "City not found"})
		return
	}

	if models.DB.Model(&cinema).Where("id = ?", id).Updates(&cinema).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func (cn *CinemaController) Delete(c *gin.Context) {

	var cinema models.Cinema

	var input struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.First(&cinema, input.ID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&cinema).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
