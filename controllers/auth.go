package controllers

import (
	"golang-shop-app-final/database"
	"golang-shop-app-final/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Router      *gin.Engine
	AuthService *services.AuthService
}

func NewAuthController(router *gin.Engine) *AuthController {
	return &AuthController{
		Router:      router,
		AuthService: services.NewAuthService(database.GetDB().DB),
	}
}

func (c *AuthController) RegisterRoutes() {
	authRouter := c.Router.Group("/auth")
	{
		authRouter.POST("/register", c.AuthService.Register)
		authRouter.POST("/login", c.AuthService.Login)
		authRouter.PATCH("/change-password", c.AuthService.ChangePassword)
		authRouter.PATCH("/:id/address", c.AuthService.UpdateAddress)
	}
}
