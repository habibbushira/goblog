package database

import (
	"log"
	"os"

	"github.com/habibbushira/goblog/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database: ")
	} else {
		log.Println("connect database successfully")
	}

	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
