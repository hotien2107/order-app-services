package main

var foodMenu = []FoodMenu{
	{
		ID:    "Iced Coffee",
		Group: "Iced Coffee",
		FoodList: []FoodItem{
			{Image: "/iced-americano.png", Name: "Iced Americano", Price: 10.25, Description: "Here is a short description for the first item. Wave Cafe Template is provided by Tooplate."},
		},
	},
	{
		ID:    "Hot Coffee",
		Group: "Hot Coffee",
		FoodList: []FoodItem{
			{Image: "/hot-americano.png", Name: "Hot Americano", Price: 8.50, Description: "Here is a short description for the first item. Wave Cafe Template is provided by Tooplate."},
			{Image: "/hot-cappuccino.png", Name: "Hot Cappuccino", Price: 9.50, Description: "Here is a short description for the first item. Wave Cafe Template is provided by Tooplate."},
		},
	},
	{
		ID:    "Fruit Juice",
		Group: "Fruit Juice",
		FoodList: []FoodItem{
			{Image: "/smoothie-1.png", Name: "Strawberry Smoothie", Price: 12.50, Description: "Here is a short description for the first item. Wave Cafe Template is provided by Tooplate."},
		},
	},
}
