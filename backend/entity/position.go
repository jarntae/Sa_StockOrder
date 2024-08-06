package entity


import (
	"time"

	"gorm.io/gorm"
)

type Position struct {

    PositionID string `gorm:"primaryKey"  json:"position_id"`

    PositionName       string `json:"position_name"`

     // Position 0..* to 1 Employee
    Employee []Employee `gorm:"foreignKey:position_id"`

    
    // เรากำหนด primaryKey  type เองเลยไม่ได้ใช้ gorm.model

    CreatedAt time.Time

    UpdatedAt time.Time

    DeletedAt gorm.DeletedAt `gorm:"index"`
}
