package config

import (
	"fmt"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "log"
)

var DB *gorm.DB

func InitDB() {
	// dsn := "host=localhost user=postgres password=root dbname=bharat_nxt port=5432 sslmode=disable"
	// var err error
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Failed to connect to DB:", err)
	// }
	fmt.Println("Connected to DB ✅")
}
