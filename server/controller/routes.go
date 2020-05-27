package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/jobs", Jobs).Methods("GET")

	log.Fatal(http.ListenAndServe(":3050", router))
}
