package models

import "time"

type Cart struct {
	CartID     uint      `gorm:"primaryKey;column:cart_id;autoIncrement"`
	CustomerID uint      `gorm:"column:customer_id;not null"`
	CartName   string    `gorm:"column:cart_name"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Customer   Customer  `gorm:"foreignKey:CustomerID;references:CustomerID"`
}

func (c *Cart) TableName() string {
	return "cart"
}
