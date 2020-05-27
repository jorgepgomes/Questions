package main

import (
	"github.com/jorgepgomes/Questions/server/api"
	"github.com/jorgepgomes/Questions/server/app"
)

func main() {

	app.InitMongo()
	api.InitRoutes()

}
