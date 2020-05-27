package main

import (
	"github.com/jorgepgomes/Questions/server/controllers"
	"github.com/jorgepgomes/Questions/server/routes"
)

func main() {

	controllers.InitMongo()
	routes.InitRoutes()

}
