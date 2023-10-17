package handlers

import (
	"github.com/gin-gonic/gin"
	"kk/config"
	"kk/models"
	"net/http"
)

// ErrorResponse represents a standardized error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a standardized success response.
type SuccessResponse struct {
	Message string `json:"message"`
}

// GetAllIndexcards retrieves all index cards.
func GetAllIndexcards(c *gin.Context) {
	var indexcards []models.Indexcard
	if err := config.DB.Find(&indexcards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, indexcards)
}

// GetIndexcardByID retrieves a specific index card by its ID.
func GetIndexcardByID(c *gin.Context) {
	var indexcard models.Indexcard
	id := c.Param("id")
	if err := config.DB.First(&indexcard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Index card not found"})
		return
	}
	c.JSON(http.StatusOK, indexcard)
}

// CreateIndexcard creates a new index card.
func CreateIndexcard(c *gin.Context) {
	var indexcard models.Indexcard
	if err := c.ShouldBindJSON(&indexcard); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	if err := config.DB.Create(&indexcard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, indexcard)
}

// DeleteIndexcard deletes a specific index card by its ID.
func DeleteIndexcard(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).Delete(&models.Indexcard{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Message: "Index card deleted successfully"})
}

// UpdateIndexcard updates a specific index card by its ID.
func UpdateIndexcard(c *gin.Context) {
	var indexcard models.Indexcard
	id := c.Param("id")
	if err := config.DB.First(&indexcard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Index card not found"})
		return
	}
	if err := c.ShouldBindJSON(&indexcard); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	if err := config.DB.Save(&indexcard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, indexcard)
}
