package handlers

import (
	"github.com/gin-gonic/gin"
	"kk/config"
	"kk/models"
	"net/http"
)

func GetAllIndexcards(c *gin.Context) {
	var indexcards []models.Indexcard

	if err := config.DB.Find(&indexcards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, indexcards)
}

func GetIndexcardByID(c *gin.Context) {
	var indexcard models.Indexcard
	id := c.Param("id")

	if err := config.DB.First(&indexcard, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, indexcard)
}

func CreateIndexcard(c *gin.Context) {
	var indexcard models.Indexcard

	if err := c.ShouldBindJSON(&indexcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&indexcard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, indexcard)
}

func DeleteIndexcard(c *gin.Context) {
	var indexcard models.Indexcard
	id := c.Param("id")

	if err := config.DB.First(&indexcard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Index card not found"})
		return
	}

	if err := config.DB.Delete(&indexcard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Index card deleted successfully"})
}

func UpdateIndexcard(c *gin.Context) {
	var indexcard models.Indexcard
	id := c.Param("id")

	if err := config.DB.First(&indexcard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Index card not found"})
		return
	}

	if err := c.ShouldBindJSON(&indexcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&indexcard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, indexcard)
}
