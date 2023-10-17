package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kk/api/routes"
	"kk/config"
	"kk/models"
)

func main() {
	config.LoadConfig()
	router := gin.Default()
	routes.InitializeRoutes(router)
	models.AutoMigrate()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("failed to start api:", err)
		return
	}
}
