package database

import (
	"golang-shop-app-final/config"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *DB
	once     sync.Once
)

type DB struct {
	DB *gorm.DB
}

func NewDB(dsn string) *DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	return &DB{DB: db}
}

func GetDB() *DB {
	config := config.GetConfig()
	once.Do(func() {
		instance = NewDB(config.Database.DSN)
	})
	return instance
}
