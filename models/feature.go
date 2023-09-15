package models

import "gorm.io/gorm"

type Feature struct {
	gorm.Model
	RouteInfo FeatureRoute `json:"routeInfo"`
}

type FeatureRoute struct {
	Name     string `json:"name"`
	Pathname string `json:"pathname"`
}

type FeatureControl struct {
	gorm.Model
}
