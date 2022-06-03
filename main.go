package main

import (
	"github.com/batthanhvan/middlewares"
	"github.com/batthanhvan/src/controllers"
	"github.com/batthanhvan/src/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDataBase()
	r := gin.Default()
	r.SetTrustedProxies(nil)

	api := r.Group("/api")
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	admin := api.Group("/admin")
	admin.Use(middlewares.JwtAuthMiddleware())
	//check token => info
	admin.GET("/user", controllers.CurrentUser)

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
