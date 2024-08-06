package entity

import (

	"gorm.io/gorm"
    
)

type Gender struct {

	gorm.Model

	GenderName string `json:"gender_name"`

	// Gender 0..* to 1 Employee
	Employee []Employee `gorm:"foreignKey:GenderID"`
}
