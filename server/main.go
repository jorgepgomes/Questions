package main

import "github.com/jorgepgomes/Questions/server/controller"

func main() {

	controller.InitMongo()
	controller.InitRoutes()

}
