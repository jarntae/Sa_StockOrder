package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {

	CategoryID   string `gorm:"primaryKey"  json:"category_id"`

	// เรากำหนด primaryKey  type เองเลยไม่ได้ใช้ gorm.model

	CreatedAt time.Time

    UpdatedAt time.Time

    DeletedAt gorm.DeletedAt `gorm:"index"`

    CategoryName string `json:"category_name"`

	//Category 1 to 0..* Product

	Product []Product  `gorm:"foreignKey:category_id"`

	
}
