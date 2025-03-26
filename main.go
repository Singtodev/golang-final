package main

import (
	"fmt"
	"golang-shop-app-final/config"
	"golang-shop-app-final/controllers"
	"golang-shop-app-final/database"
	"golang-shop-app-final/models"

	"github.com/gin-gonic/gin"
)

func main() {
	instance := database.GetDB()
	if instance == nil {
		fmt.Println("Database connection failed")
		return
	}
	gin.SetMode(config.GetConfig().Server.Mode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		customers := []models.Customer{}
		instance.DB.Find(&customers)
		fmt.Println(customers)
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	authController := controllers.NewAuthController(router)
	authController.RegisterRoutes()
	router.Run(":" + config.GetConfig().Server.Port)
}
