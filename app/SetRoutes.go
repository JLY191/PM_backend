package app

import (
	"PM_backend/app/controller"
	"PM_backend/app/middleware"
)

func addRoutes() {
	G.GET("/ping", middleware.Cors(), controller.PingHandler)

	userGroup := G.Group("/user", middleware.Cors())
	userGroup.POST("/register", controller.RegisterHandler)
	userGroup.POST("/login", controller.LoginHandler)
	userGroup.GET("/logout", controller.LogoutHandler)

	siteGroup := G.Group("/site", middleware.Cors(), middleware.Auth())
	siteGroup.POST("/search", controller.SearchHandler)

	remarkGroup := G.Group("/remark", middleware.Cors(), middleware.Auth())
	remarkGroup.POST("/add", controller.AddRemarkHandler)
}
