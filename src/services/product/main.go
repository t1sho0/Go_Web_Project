package main

import (
	"fmt"
	"log"
	"os"
	"product_service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST") 
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	err = db.AutoMigrate(
		&models.Brand{},
		&models.Category{},
		&models.Product{},
	)

	if err != nil {
		log.Fatal("Ошибка миграции: ", err)
	}
	fmt.Println("База данных успешно проинициализированна и мигрированна.")

	return db
}

func main() {
	db := InitDB()
	r := gin.Default()

	
}
