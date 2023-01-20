package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/naumaan12/go-api/models" 
  "github.com/naumaan12/go-api/controllers"
)

func main() {
  r := gin.Default()
  
  models.ConnectDatabase()
  r.POST("/users", controllers.CreateUser)
  r.PATCH("/users/:id", controllers.UpdateUser) 
  r.DELETE("/users/:id", controllers.DeleteUser)
  
  
  r.GET("/users", controllers.FindUsers)
  r.GET("/users/:id", controllers.FindUser) 

  
  r.GET("/", func(c *gin.Context) {
  	c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.Run()
}
