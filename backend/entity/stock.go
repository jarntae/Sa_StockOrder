package entity

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model

	StockID       string      `json:"stock_id"`

	StockName     string      `json:"stock_name"`

	CategoryID    uint        `json:"category_id"` 

	Category      *Category   `gorm:"foreignKey:CategoryID"` 
	
	QuantityInStock int       `json:"quantity"`
	
	LastUpdated   time.Time   `json:"last_updated"`
	
	UpdatedByEmployeeID uint   `json:"updated_by_employee_id"` 

	UpdatedByEmployee   *Employee `gorm:"foreignKey:UpdatedByEmployeeID"` 
}
