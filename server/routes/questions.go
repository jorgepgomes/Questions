package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgepgomes/Questions/server/controllers"
)

func SetQuestionsRouters(router *mux.Router) *mux.Router {

	router.HandleFunc("/api/questions/create", controllers.CreateQuestion).Methods("POST")
	router.HandleFunc("/api/questions/answers", controllers.AnswerQuestion).Methods("POST")

	return router
}
