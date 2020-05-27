package controller

import (
	"fmt"

	"github.com/jorgepgomes/Questions/server/model"
	"github.com/jorgepgomes/Questions/server/nosql"
	"gopkg.in/mgo.v2"
)

type App struct {
}

type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func InitMongo() {

	model.ReadConfig()

	fmt.Println("Init MongoDB")
	session := nosql.InitialiseMongo()
	mongoStore.session = session
}
