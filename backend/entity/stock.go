package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model

	Quantity uint `json:"stock_quantity"`

	Price float64 `json:"stock_price"`

	DateIn time.Time `json:"stock_date_in"`

	ExpirationDate time.Time `json:"stock_expiration_date"`

	ProductID *string

	Product Product `gorm:"foriegnKey:product_id"`

	EmployeeID *string

	Employee Employee `gorm:"foriegnKey:EmployeeID"`

	SupplierID *uint

	Supplier Supplier `gorm:"foriegnKey:SupplierID"`

}
