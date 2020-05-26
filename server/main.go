package main

import (
	"fmt"
	"net/http"

	"github.com/jorgepgomes/Questions/server/common"
	"github.com/jorgepgomes/Questions/server/routers"
)

func main() {

	common.StartUp()
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	fmt.Println("Listening...")
	server.ListenAndServe()
}
