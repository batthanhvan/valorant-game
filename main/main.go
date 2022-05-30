package main

import (
	"github.com/batthanhvan/src/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	reportGroup := r.Group("/reports")
	reportGroup.GET("/show", controllers.HandleShowAllReports)
	reportGroup.GET("/search", controllers.HandleGetReportByUsername)

	playerGroup := r.Group("/players")
	playerGroup.GET("/:username", controllers.HandleGetByUserName)
	// playerGroup.POST("/restrict/:username", controllers.HandlePostRestrictPlayerTime)
	playerGroup.POST("/:username/:playername/:tagline", controllers.HandlePostModifyUser)

	matchGroup := r.Group("/matches")
	matchGroup.GET("/search", controllers.HandleGetByMatchID)

	r.Run("localhost:8080")
}
