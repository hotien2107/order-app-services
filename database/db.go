package database

import (
	"order-app/web-service-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "user=postgres password=Speedattack2107 host=db.jrpbyyrgtuajhoswvywd.supabase.co port=5432 dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.FoodGroup{})
	db.AutoMigrate(&models.FoodItem{})

	DB = db
}
