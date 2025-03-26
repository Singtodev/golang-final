package controllers

import (
	"golang-shop-app-final/database"
	"golang-shop-app-final/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Router         *gin.Engine
	ProductService *services.ProductService
}

func NewProductController(router *gin.Engine) *ProductController {
	return &ProductController{
		Router:         router,
		ProductService: services.NewProductService(database.GetDB().DB),
	}
}

func (c *ProductController) RegisterRoutes() {
	productRouter := c.Router.Group("/products")
	{
		productRouter.GET("", c.ProductService.GetProducts)
		productRouter.GET("/search", c.ProductService.SearchProducts)
	}
}
