package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naumaan12/go-api/models"
)

type UpdateUserInput struct {
	Name         string `json:"name"`
	MobileNumber int    `json:"mobileNumber"`
	Age          int    `json:"age"`
	Password     string `json:"password"`
}

func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.USER
	if err := models.DB.Where("email = ?", c.Param("email")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
