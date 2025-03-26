package models

import "time"

type CartItem struct {
	CartItemID uint      `gorm:"primaryKey;column:cart_item_id;autoIncrement"`
	CartID     uint      `gorm:"column:cart_id;not null"`
	ProductID  uint      `gorm:"column:product_id;not null"`
	Quantity   int       `gorm:"column:quantity;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Cart       Cart      `gorm:"foreignKey:CartID;references:CartID"`
	Product    Product   `gorm:"foreignKey:ProductID;references:ProductID"`
}

func (c *CartItem) TableName() string {
	return "cart_item"
}
