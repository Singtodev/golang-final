package controllers

import (
	"golang-shop-app-final/database"
	"golang-shop-app-final/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Router          *gin.Engine
	CustomerService *services.CustomerService
}

func NewCustomerController(router *gin.Engine) *CustomerController {
	return &CustomerController{
		Router:          router,
		CustomerService: services.NewCustomerService(database.GetDB().DB),
	}
}

func (c *CustomerController) RegisterRoutes() {
	customerRouter := c.Router.Group("/customers")
	{
		customerRouter.GET("/:id/cart", c.CustomerService.GetAllCart)
		customerRouter.POST("/:id/cart/add", c.CustomerService.AddProductToCart)
	}
}
