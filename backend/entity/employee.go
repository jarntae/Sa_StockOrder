package entity

import (

	"gorm.io/gorm"
)

type Employee struct {

	gorm.Model

    FirstName    string    `json:"first_name"`

    LastName     string    `json:"last_name"`

    Username     string    `json:"username"`

    Password     string    `json:"password"`
	
	GenderID     uint   `json:"gender_id"`

	PositionID   uint   `json:"position_id"`

	//Employee 1 to 1..* Product

	Product		[]Product	 `gorm:"foreignKey:EmployeeID"`

	//Employee 1 to 1..* Stock

	Stock []Stock `gorm:"foreignKey:EmployeeID"`

}