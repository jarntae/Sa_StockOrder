package entity


import "gorm.io/gorm"


type Employee struct {

	gorm.Model

	EmployeeID string `json:"employee_id"`

	FirstName  string `json:"firstname"`

	LastName   string `json:"lastname"`

	Gender     string `json:"gender"`

	Position   string `json:"position"`

	Username   string `json:"username"`
	
	Password   string `json:"password"`
}