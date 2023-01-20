package models

type USER struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	Email        string `json:"email"  `
	MobileNumber int    `json:"mobileNumber"`
	Age          int    `json:"age"`
	Password     string `json:"password"`
}
