package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jorgepgomes/Questions/server/model"
)

func ReadConfig() {
	file, _ := ioutil.ReadFile("../config.json")

	data := model.Config{}

	_ = json.Unmarshal([]byte(file), &data)
	model.Cfg = &data
}
