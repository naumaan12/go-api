package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naumaan12/go-api/models"
)


func FindUsers(c *gin.Context) {
	var users []models.USER
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
