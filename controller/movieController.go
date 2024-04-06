package controller

import (
	"net/http"
	"ticketing_movies_API_GO/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieController struct{}

func (mc *MovieController) Index(c *gin.Context) {
	var movies []models.Movie

	if err := models.DB.Find(&movies).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch movies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Movies": movies})

}

func (mc *MovieController) Show(c *gin.Context) {
	var movie models.Movie

	id := c.Param("id")

	if err := models.DB.First(&movie, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Movie": movie})
}

func (mc *MovieController) Create(c *gin.Context) {

	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&movie)
	c.JSON(http.StatusOK, gin.H{"Movie": movie})

}

func (mc *MovieController) Update(c *gin.Context) {
	var movie models.Movie

	id := c.Param("id")

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&movie).Where("id = ?", id).Updates(&movie).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func (mc *MovieController) Delete(c *gin.Context) {

	var movie models.Movie

	var input struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := models.DB.First(&movie, input.ID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	if err := models.DB.Delete(&movie).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
