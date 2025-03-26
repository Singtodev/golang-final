package services

import (
	"golang-shop-app-final/dtos"
	"golang-shop-app-final/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (s *AuthService) Register(c *gin.Context) {
	customer := dtos.CustomerDTO{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCustomer models.Customer
	if result := s.DB.Where("email = ?", customer.Email).First(&existingCustomer); result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	customerModel := models.Customer{}
	copier.Copy(&customerModel, &customer)
	customerModel.Password = string(hashedPassword)

	if err := s.DB.Create(&customerModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer created successfully"})
}

func (s *AuthService) Login(c *gin.Context) {
	loginData := dtos.CustomerLoginDTO{}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer
	if result := s.DB.Where("email = ?", loginData.Email).First(&customer); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (s *AuthService) ChangePassword(c *gin.Context) {
	customer := dtos.CustomerChangePasswordDTO{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCustomer models.Customer
	if result := s.DB.Where("email = ?", customer.Email).First(&existingCustomer); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingCustomer.Password), []byte(customer.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid old password"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	existingCustomer.Password = string(hashedPassword)

	if err := s.DB.Save(&existingCustomer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func (s *AuthService) UpdateAddress(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	var existingCustomer models.Customer
	if result := s.DB.Where("customer_id = ?", id).First(&existingCustomer); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	customer := dtos.CustomerDTO{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingCustomer.Address = customer.Address

	if err := s.DB.Save(&existingCustomer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Address updated successfully",
		"customer": existingCustomer,
	})
}
