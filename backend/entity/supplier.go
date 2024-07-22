package entity

import "gorm.io/gorm"

type Supplier struct {

	gorm.Model

	SupplierID   int    `json:"supplier_id"`
	
	SupplierName string `json:"supplier_name"`

	Phone        string `json:"phone"`

	Email        string `json:"email"`
	
	Address      string `json:"address"`
}