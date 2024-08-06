package entity

import (

	"gorm.io/gorm"
)

type Position struct {

    gorm.Model

    PositionName       string `json:"position_name"`

     // Position 0..* to 1 Employee
    Employee []Employee `gorm:"foreignKey:PositionID"`

}
