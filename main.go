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
  r.PATCH("/users/:email", controllers.UpdateUser) 
  r.DELETE("/users/:email", controllers.DeleteUser)
  
  
  r.GET("/users", controllers.FindUsers)
  r.GET("/users/:email", controllers.FindUser) 

  
  r.GET("/", func(c *gin.Context) {
  	c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.Run()
}
