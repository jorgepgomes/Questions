package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jorgepgomes/Questions/server/model"
	"gopkg.in/mgo.v2/bson"
)

func (a *App) JobsGetHandler() string {

	config := model.Cfg
	mongo := config.Server.MongoDB
	database := mongo.Database
	collection := mongo.Collection

	col := mongoStore.session.DB(database).C(collection)

	results := []model.Job{}
	col.Find(bson.M{"title": bson.RegEx{"", ""}}).All(&results)
	jsonString, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	return string(jsonString)

}

func JobsPostHandler(w http.ResponseWriter, r *http.Request) {

	config := model.Cfg
	mongo := config.Server.MongoDB
	database := mongo.Database
	collection := mongo.Collection

	col := mongoStore.session.DB(database).C(collection)

	//Retrieve body from http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//Save data into Job struct
	var _job model.Job
	err = json.Unmarshal(b, &_job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Insert job into MongoDB
	err = col.Insert(_job)
	if err != nil {
		panic(err)
	}

	//Convert job struct into json
	jsonString, err := json.Marshal(_job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back data as response
	w.Write(jsonString)

}
