package app

import (
	"PM_backend/app/controller"
)

func addRoutes() {
	G.GET("/ping", controller.PingHandler)
	userGroup := G.Group("/user")
	userGroup.POST("/register", controller.RegisterHandler)
	userGroup.POST("/login", controller.LoginHandler)
	userGroup.GET("/logout", controller.LogoutHandler)
}
