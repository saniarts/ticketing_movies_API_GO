package controller

import (
	"net/http"
	"ticketing_movies_API_GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CityController struct{}

func (cc *CityController) Index(c *gin.Context) {
	var cities []models.City

	if err := models.DB.Find(&cities).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch cities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cities": cities})

}

func (cc *CityController) Show(c *gin.Context) {
	var city models.City

	id := c.Param("id")

	if err := models.DB.First(&city, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"city": city})
}

func (cc *CityController) Create(c *gin.Context) {

	var city models.City

	if err := c.ShouldBindJSON(&city); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.Create(&city).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"city": city})

}

func (cc *CityController) Update(c *gin.Context) {
	var city models.City

	id := c.Param("id")

	if err := c.ShouldBindJSON(&city); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&city).Where("id = ?", id).Updates(&city).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func (cc *CityController) Delete(c *gin.Context) {

	var city models.City

	var input struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.First(&city, input.ID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&city).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
