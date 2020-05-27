package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() {
	fmt.Println("Init Route")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/jobs", Jobs).Methods("GET")

	http.ListenAndServe(":3050", router)
	fmt.Println("Listen: 3050")
}
