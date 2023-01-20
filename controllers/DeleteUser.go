
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naumaan12/go-api/models"
)


func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.USER
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&user)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }
