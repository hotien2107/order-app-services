package models

import "gorm.io/gorm"

type FoodGroup struct {
	gorm.Model
	Name     string     `json:"name"`
	FoodList []FoodItem `json:"foodList"`
}

type FoodItem struct {
	gorm.Model
	FoodGroupID uint
	Group       FoodGroup `json:"group" gorm:"foreignKey:FoodGroupID"`
	Image       string    `json:"image"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	Description string    `json:"desc"`
}
