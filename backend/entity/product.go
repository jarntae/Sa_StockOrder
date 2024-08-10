package entity

import (

	"gorm.io/gorm"
)

type Product struct {

	gorm.Model

	Product_Code_ID   string `json:"product_code_id"`

    ProductName string `json:"product_name"`

    CategoryID  *string 

	Category	Category	 `gorm:"foriegnKey:CategoryID"`

    EmployeeID  *string    

	Employee	Employee	 `gorm:"foriegnKey:EmployeeID"`

	//Product 1 to 1..* Stock

	Stock	[]Stock `gorm:"foreignKey:ProductID"`

}