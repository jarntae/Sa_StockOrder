package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	StockID uint `gorm:"primaryKey" json:"stock_id"`

	Quantity uint `json:"quantity"`

	Price float64 `json:"price"`

	DateIn time.Time `json:"date_in"`

	ExpirationDate time.Time `json:"expiration_date"`

	ProductID *string

	Product Product `gorm:"foriegnKey:product_id"`

	EmployeeID *string

	Employee Employee `gorm:"foriegnKey:employee_id"`

	SupplierID *uint

	Supplier Supplier `gorm:"foriegnKey:supplier_id"`

	// เรากำหนด primaryKey  type เองเลยไม่ได้ใช้ gorm.model

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}
