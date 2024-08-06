package entity

import (
        "gorm.io/gorm"

        "time"
)

type Gender struct {

	GenderID string `gorm:"primaryKey"  json:"gender_id"`

	GenderName string `json:"gender_name"`

    CreatedAt time.Time

    UpdatedAt time.Time

    DeletedAt gorm.DeletedAt `gorm:"index"`

	// Gender 0..* to 1 Employee
	Employee []Employee `gorm:"foreignKey:gender_id"`
}
