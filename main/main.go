package main

import (
	"github.com/batthanhvan/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	playerGroup := r.Group("/player")
	playerGroup.GET("/search", controllers.HandleGetByUserName)
	r.Run("localhost:8080")
}
