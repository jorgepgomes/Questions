package api

import (
	"net/http"

	"github.com/jorgepgomes/Questions/server/model"
)

func Jobs(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(model.ToJson("Ok")))
}
