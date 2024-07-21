package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	CategoryID string `json:"category_id"`

	CategoryName string `json:"category_name"`
}
