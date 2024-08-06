package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {

	ProductID   string `gorm:"primaryKey"  json:"product_id"`

	// เรากำหนด primaryKey  type เองเลยไม่ได้ใช้ gorm.model

	CreatedAt time.Time

    UpdatedAt time.Time

    DeletedAt gorm.DeletedAt `gorm:"index"`


    ProductName string `json:"product_name"`

    CategoryID  *string 

	Category	Category	 `gorm:"foriegnKey:category_id"`

    EmployeeID  *string    

	Employee	Employee	 `gorm:"foriegnKey:EmployeeID"`

	Stock	[]Stock `gorm:"foreignKey:product_id"`

}