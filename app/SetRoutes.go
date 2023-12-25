package app

import "PM_backend/app/controller"

func addRoutes() {
	G.GET("/ping", controller.PingHandler)
}
