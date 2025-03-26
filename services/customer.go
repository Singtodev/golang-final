package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerService struct {
	DB *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{
		DB: db,
	}
}

func (s *CustomerService) UpdateCustomer(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Update Customer"})
}
