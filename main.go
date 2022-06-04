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

	// api := r.Group("/api")
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/user", controllers.CurrentUser)

	admin := r.Group("/admin")
	// admin.Use(middlewares.AuthenticateRole("admin"))

	reportGroup := admin.Group("/reports")
	reportGroup.Use(middlewares.AuthenticateRole("admin"))
	reportGroup.GET("/show", controllers.HandleShowAllReports)
	reportGroup.GET("/search", controllers.HandleGetReportByUsername)

	playerGroup := r.Group("/players")
	playerGroup.GET("/:username", controllers.HandleGetByUserName)
	// playerGroup.POST("/restrict/:username", controllers.HandlePostRestrictPlayerTime)
	playerGroup.POST("/modify/:playername/:tagline", controllers.HandlePostModifyUser)

	matchGroup := r.Group("/matches")
	matchGroup.GET("/search", controllers.HandleGetByMatchID)

	r.Run("localhost:8080")
}
