package routes

import (
	"github.com/gin-gonic/gin"
	"kk/api/handlers"
)

func InitializeRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Health check endpoint
		api.GET("/health", handlers.HealthHandler)

		// Indexcard
		api.GET("/indexcards", handlers.GetAllIndexcards)
		api.GET("/indexcard", handlers.GetIndexcardByID)
		api.POST("/indexcard", handlers.CreateIndexcard)
		api.DELETE("/indexcard", handlers.DeleteIndexcard)
		api.PATCH("/indexcard", handlers.UpdateIndexcard)
	}
}
