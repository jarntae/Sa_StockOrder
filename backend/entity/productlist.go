package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProductList struct {

	gorm.Model

	StockID           string    `json:"stock_id"`

	StockName         string    `json:"stock_name"`

	CategoryID        string    `json:"id_category"`

	Category          *Category `gorm:"foreignKey:id_category" json:"category_id"`

	LastUpdated       time.Time `json:"last_updated"`

	EmployeeUpdatedID uint       `json:"employee_updated_id"`

	EmployeeUpdated   *Employee  `gorm:"foreignKey:employee_updated_id" json:"employee_id"`
}