package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() {
	fmt.Println("Init Route")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/jobs", JobsGetHandler).Methods("GET")

	fmt.Println("Listen in: 3050")
	http.ListenAndServe(":3050", router)
}
