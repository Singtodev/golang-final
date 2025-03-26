package models

import "time"

type Product struct {
	ProductID     uint      `gorm:"primaryKey;column:product_id;autoIncrement"`
	ProductName   string    `gorm:"column:product_name;not null"`
	Description   string    `gorm:"column:description;type:text"`
	Price         float64   `gorm:"column:price;type:decimal(10,2);not null"`
	StockQuantity int       `gorm:"column:stock_quantity;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (c *Product) TableName() string {
	return "product"
}
