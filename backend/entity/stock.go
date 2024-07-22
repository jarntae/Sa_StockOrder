package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {

	gorm.Model

	StockID           string     `json:"stock_id"`

	SupplierID        uint        `json:"supplier_id"`

	Supplier          *Supplier  `gorm:"foreignKey:SupplierID"`

	Quantity          int        `json:"quantity"`

	Price             float64    `json:"price" gorm:"type:decimal(7,2)"`

	DateIn            time.Time  `json:"date_in"`

	ExpirationDate    time.Time  `json:"expiration_date"`

	EmployeeUpdatedID uint        `json:"employee_updated_id"`

	EmployeeUpdated   *Employee  `gorm:"foreignKey:employee_updated_id" json:"employee_id"`
}
