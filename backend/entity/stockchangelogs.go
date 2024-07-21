package entity

import (
	"time"

	"gorm.io/gorm"
)

type StockChangeLog struct {
	gorm.Model
	LogID              int       `gorm:"primaryKey"`

	StockID            uint      `json:"stock_id"`

	Stock              *Stock    `gorm:"foreignKey:StockID"`

	QuantityChanged    int       `json:"quantity_changed"`

	ChangeType         string    `json:"change_type"`

	ChangeDate         time.Time `json:"change_date"`

	ChangedByEmployeeID uint     `json:"changed_by_employee_id"`
	
	ChangedByEmployee  *Employee `gorm:"foreignKey:ChangedByEmployeeID"`
}