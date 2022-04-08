package main

import (
	"github.com/batthanhvan/src/api"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", HomePage)
	playerGroup := r.Group("/player")
	playerGroup.GET("/search", api.HandleGetByUserName)
	r.Run("localhost:8080")
}
