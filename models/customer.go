package models

import "time"

type Customer struct {
	CustomerID  uint      `gorm:"primaryKey;autoIncrement"`
	FirstName   string    `gorm:"type:varchar(255);not null"`
	LastName    string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"type:varchar(255);unique;not null"`
	PhoneNumber string    `gorm:"type:varchar(20)"`
	Address     string    `gorm:"type:varchar(255)"`
	Password    string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (c *Customer) TableName() string {
	return "customer"
}
