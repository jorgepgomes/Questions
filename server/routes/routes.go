package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitRoutes() {
	fmt.Println("Init Route")
	router := mux.NewRouter().StrictSlash(true)
	router = SetQuestionsRouters(router)
	handler := cors.Default().Handler(router)

	fmt.Println("Listen in: 3050")
	http.ListenAndServe(":3050", handler)
}
