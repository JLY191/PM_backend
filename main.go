package main

import (
	"PM_backend/app"
	"PM_backend/model"
)

func main() {
	model.InitDB()
	app.InitWeb()

}
