package controller

package app

import (
	"fmt"

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

	ReadConfig()

	fmt.Println("Init MongoDB")
	session := nosql.InitialiseMongo()
	mongoStore.session = session
}
