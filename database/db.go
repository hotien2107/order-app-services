package database

import (
	"order-app/web-service-gin/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	pass := os.Getenv("DB_PASS")
	dsn := "postgres://postgres:" + pass + "@db.ozhtmcffuhdwetgektfq.supabase.co:6543/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.FoodGroup{})
	db.AutoMigrate(&models.FoodItem{})

	DB = db
}
