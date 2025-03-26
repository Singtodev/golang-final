package services

import (
	"golang-shop-app-final/dtos"
	"golang-shop-app-final/models"
	"net/http"
	"strconv"

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

func (s *CustomerService) AddProductToCart(c *gin.Context) {
	customerID := c.Param("id")
	customerAddItemDto := dtos.CustomerAddItemDto{}

	if err := c.ShouldBindJSON(&customerAddItemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customer{}
	if err := s.DB.First(&customer, customerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	product := models.Product{}
	if err := s.DB.First(&product, customerAddItemDto.ProductID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if customerAddItemDto.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than 0"})
		return
	}

	cart := models.Cart{}
	if err := s.DB.First(&cart, "customer_id = ? AND cart_name = ?", customerID, customerAddItemDto.CartName).Error; err != nil {
		customerIDUint := uint(0)
		if id, err := strconv.ParseUint(customerID, 10, 32); err == nil {
			customerIDUint = uint(id)
		}
		cart = models.Cart{
			CustomerID: customerIDUint,
			CartName:   customerAddItemDto.CartName,
		}
		if err := s.DB.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
			return
		}
	}

	// Check if item already exists in cart
	existingItem := models.CartItem{}
	if err := s.DB.First(&existingItem, "cart_id = ? AND product_id = ?", cart.CartID, customerAddItemDto.ProductID).Error; err == nil {
		// Item exists, increment quantity
		existingItem.Quantity += customerAddItemDto.Quantity
		if err := s.DB.Save(&existingItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item quantity"})
			return
		}
	} else {
		// Item doesn't exist, create new cart item
		cartItem := models.CartItem{
			CartID:    cart.CartID,
			ProductID: customerAddItemDto.ProductID,
			Quantity:  customerAddItemDto.Quantity,
		}
		if err := s.DB.Create(&cartItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully"})
}

func (s *CustomerService) GetAllCart(c *gin.Context) {
	customerID := c.Param("id")

	// ตรวจสอบว่า customer มีอยู่จริงหรือไม่
	var customer models.Customer
	if err := s.DB.First(&customer, customerID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	var carts []models.Cart
	if err := s.DB.Where("customer_id = ?", customerID).Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch carts"})
		return
	}

	// ถ้าไม่มีตะกร้าสินค้าเลย
	if len(carts) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No carts found",
			"carts":   []interface{}{},
		})
		return
	}

	type CartItemDetail struct {
		CartItemID  uint    `json:"cart_item_id"`
		ProductID   uint    `json:"product_id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Quantity    int     `json:"quantity"`
		Total       float64 `json:"total"`
	}

	type CartDetail struct {
		CartID    uint             `json:"cart_id"`
		CartName  string           `json:"cart_name"`
		Items     []CartItemDetail `json:"items"`
		CartTotal float64          `json:"cart_total"`
	}

	var response []CartDetail

	for _, cart := range carts {
		var cartItems []models.CartItem
		if err := s.DB.Preload("Product").Where("cart_id = ?", cart.CartID).Find(&cartItems).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
			return
		}

		cartDetail := CartDetail{
			CartID:    cart.CartID,
			CartName:  cart.CartName,
			Items:     []CartItemDetail{}, // กำหนดค่าเริ่มต้นเป็น array ว่าง
			CartTotal: 0,
		}

		// ถ้ามีสินค้าในตะกร้า
		if len(cartItems) > 0 {
			for _, item := range cartItems {
				itemDetail := CartItemDetail{
					CartItemID:  item.CartItemID,
					ProductID:   item.ProductID,
					Name:        item.Product.ProductName,
					Description: item.Product.Description,
					Price:       item.Product.Price,
					Quantity:    item.Quantity,
					Total:       float64(item.Quantity) * item.Product.Price,
				}
				cartDetail.Items = append(cartDetail.Items, itemDetail)
				cartDetail.CartTotal += itemDetail.Total
			}
		}

		response = append(response, cartDetail)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Carts retrieved successfully",
		"carts":   response,
	})
}
