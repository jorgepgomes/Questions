package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jorgepgomes/Questions/server/model"
	"gopkg.in/mgo.v2/bson"
)

func JobsGetHandler(w http.ResponseWriter, r *http.Request) {

	database := model.Cfg.Server.MongoDB.Database
	collection := model.Cfg.Server.MongoDB.Collection

	col := mongoStore.session.DB(database).C(collection)

	// results := []model.Job{}

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	x := col.Find(bson.M{})
	fmt.Println(">>>> ", x)
	// col.Find(bson.M{"title": bson.RegEx{"", ""}}).All(&results)
	// jsonString, err := json.Marshal(results)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(">>> ", string(jsonString))
	// fmt.Fprint(w, string(jsonString))

}

func Add(w http.ResponseWriter, r *http.Request) {

	database := model.Cfg.Server.MongoDB.Database
	collection := model.Cfg.Server.MongoDB.Collection

	col := mongoStore.session.DB(database).C(collection)
	// col := session.DB(database).C(collection)

	err := col.Insert(&model.Job{"DevOps Engineer", "Should be familiar with Jenkins Pipeline", "Company XYZ", "$10,000"},
		&model.Job{"Senior Software Engineer", "Should be familiar with golang", "Company XYZ", "$12,000"})

	if err != nil {
		log.Fatal(err)
	}

	count, err := col.Count()
	if err != nil {
		panic(err)
	}
	res := fmt.Sprintf("Messages count: %d", count)

	fmt.Fprint(w, res)

}

func JobsPostHandler(w http.ResponseWriter, r *http.Request) {

	database := model.Cfg.Server.MongoDB.Database
	collection := model.Cfg.Server.MongoDB.Collection

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
