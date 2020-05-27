package app

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jorgepgomes/Questions/server/model"
)

func (a *App) ReadConfig() model.ConfigServer {
	file, _ := ioutil.ReadFile("../config.json")

	data := model.ConfigServer{}

	_ = json.Unmarshal([]byte(file), &data)

	return data

}
