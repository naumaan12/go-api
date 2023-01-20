package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	MobileNumber string `json:"mobileNumber" binding:"required"`
	Age          int    `json:"age" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

// GET /users
// Get all users
func FindBooks(c *gin.Context) {
	var users []models.USER
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.USER{Name: input.Name, Email: input.Email, Password: input.Password, MobileNumber: input.MobileNumber, Age: input.Age}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func FindBook(c *gin.Context) { // Get model if exist
	var user models.USER

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
