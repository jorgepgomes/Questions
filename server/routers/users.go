package routers

import (
	"github.com/gorilla/mux"
	"github.com/jorgepgomes/Questions/server/controllers"
)

func SetUsersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	return router
}
