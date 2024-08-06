package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {

	EmployeeID   string    `gorm:"primaryKey"  json:"employee_id"`

    FirstName    string    `json:"first_name"`

    LastName     string    `json:"last_name"`

    Username     string    `json:"username"`

    Password     string    `json:"password"`

    RegisterDate time.Time `json:"register_date"`

	
	GenderID     *string   

	Gender		Gender     `gorm:"foriegnKey:gender_id"`

	PositionID   *string   

	Position	Position	`gorm:"foriegnKey:position_id"`

	Product		[]Product	 `gorm:"foreignKey:employee_id"`

	Stock []Stock `gorm:"foreignKey:employee_id"`

	// เรากำหนด primaryKey  type เองเลยไม่ได้ใช้ gorm.model

	CreatedAt time.Time

    UpdatedAt time.Time

    DeletedAt gorm.DeletedAt `gorm:"index"`

}