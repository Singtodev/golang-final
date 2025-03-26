package services

import (
	"golang-shop-app-final/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{
		DB: db,
	}
}

func (s *ProductService) GetProducts(c *gin.Context) {
	products := []models.Product{}
	if err := s.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get Products", "products": products})
}

func (s *ProductService) SearchProducts(c *gin.Context) {
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	search := c.Query("search")

	products := []models.Product{}
	query := s.DB

	if minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}

	if maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}

	if search != "" {
		query = query.Where("product_name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Search Products", "products": products})
}
