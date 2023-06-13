package main

type FoodItem struct {
	Image       string  `json:"image"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"desc"`
}

type FoodMenu struct {
	ID       string     `json:"id"`
	Group    string     `json:"group"`
	FoodList []FoodItem `json:"foodList"`
}
