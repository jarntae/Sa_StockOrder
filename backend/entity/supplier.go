package entity

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	SupplierID uint `gorm:"primaryKey" json:"supplier_id"`

	SupplierName string `json:"supplier_name"`

	Phone string `json:"phone"`

	Email string `json:"email"`

	Address string `json:"address"`

	Stock []Stock `gorm:"foreignKey:supplier_id"`

	// เรากำหนด primaryKey  type เองเลยไม่ได้ใช้ gorm.model

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}
