package main

import (
	"github.com/batthanhvan/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	reportGroup := r.Group("/reports")
	reportGroup.GET("/show", controllers.HandleShowAllReports)
	reportGroup.GET("/search", controllers.HandleGetReportByUsername)

	playerGroup := r.Group("/players")
	playerGroup.GET("/search", controllers.HandleGetByUserName)

	matchGroup := r.Group("/matches")
	matchGroup.GET("/search", controllers.HandleGetByMatchID)

	r.Run("localhost:8080")
}
