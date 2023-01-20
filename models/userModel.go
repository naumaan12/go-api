package models

type USER struct {
	ID           uint   `json:"id" `
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"primary_key" `
	MobileNumber int    `json:"mobileNumber"`
	Age          int    `json:"age"`
	Password     string `json:"password"`
}
