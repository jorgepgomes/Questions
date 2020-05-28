package nosql

import (
	"fmt"
	"time"

	"github.com/jorgepgomes/Questions/server/model"
	"gopkg.in/mgo.v2"
)

func InitialiseMongo() (session *mgo.Session) {

	config := model.Cfg
	fmt.Println(">>> ", config)
	mongo := config.Server.MongoDB

	info := &mgo.DialInfo{
		Addrs:    []string{mongo.Hosts},
		Timeout:  60 * time.Second,
		Database: mongo.Database,
		Username: mongo.Username,
		Password: mongo.Password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return

}
