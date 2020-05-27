package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() {
	fmt.Println("Init Route")
	router := mux.NewRouter().StrictSlash(true)

	router = SetQuestionsRouters(router)

	fmt.Println("Listen in: 3050")
	http.ListenAndServe(":3050", router)
}
