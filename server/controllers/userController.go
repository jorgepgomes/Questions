package controllers

import (
	"encoding/json"
	"net/http"

	common "github.com/jorgepgomes/Questions/server/common"
	data "github.com/jorgepgomes/Questions/server/data"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("questions")
	repo := &data.UserRepository{c}

	questions := repo.GetAll()
	j, err := json.Marshal(UsersResource{Data: questions})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var dataResource UserResource
// 	// Decode the incoming User json
// 	err := json.NewDecoder(r.Body).Decode(&dataResource)
// 	if err != nil {
// 		common.DisplayAppError(w, err, "Invalid User data", 500)
// 		return
// 	}
// 	user := &dataResource.Data
// 	// Create new context
// 	context := NewContext()
// 	defer context.Close()
// 	c := context.DbCollection("questions")
// 	// Create User
// 	repo := &data.UserRepository{c}
// 	repo.Create(user)
// 	// Create response data
// 	j, err := json.Marshal(dataResource)
// 	if err != nil {
// 		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
// 		return
// 	}
// 	// Send response back
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(j)
// }

// // Handler for HTTP Delete - "/users/{id}"
// // Delete a User document by id
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	// Get id from incoming url
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	// Create new context
// 	context := NewContext()
// 	defer context.Close()
// 	c := context.DbCollection("users")

// 	// Remove user by id
// 	repo := &data.UserRepository{c}
// 	err := repo.Delete(id)
// 	if err != nil {
// 		if err == mgo.ErrNotFound {
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		} else {
// 			common.DisplayAppError(w, err, "An unexpected error ahs occurred", 500)
// 			return
// 		}
// 	}

// 	// Send response back
// 	w.WriteHeader(http.StatusNoContent)
// }
