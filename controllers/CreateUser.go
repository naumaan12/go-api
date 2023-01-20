package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naumaan12/go-api/models"
)

type CreateUserInput struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	MobileNumber int    `json:"mobileNumber" binding:"required"`
	Age          int    `json:"age" binding:"required"`
	Password     string `json:"password" binding:"required"`
}


func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	fmt.Println(input)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users []models.USER
	models.DB.Find(&users)
	for i := 0; i < len(users) ; i++ {
		userData := users[i]
		if userData.Email == input.Email {
			c.JSON(http.StatusBadRequest , gin.H{"Error": "User Already Exists"})
			return
		}
	}

	// Create user
	user := models.USER{Name: input.Name, Email: input.Email, Password: input.Password, MobileNumber: input.MobileNumber, Age: input.Age}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
