package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgepgomes/Questions/server/controllers"
)

func SetQuestionsRouters(router *mux.Router) *mux.Router {

	router.HandleFunc("/jobs", controllers.JobsGetHandler).Methods("GET")
	router.HandleFunc("/api/add", controllers.Add).Methods("GET")

	return router
}
