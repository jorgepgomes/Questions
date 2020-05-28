package model

import (
	"encoding/json"
	"io/ioutil"
)

var Cfg *Config = &Config{}

type Config struct {
	Server server `json:"server"`
}

type server struct {
	Port    string `json:"port"`
	MongoDB mongo  `json:"mongodb"`
}

type mongo struct {
	Hosts      string `json:"hosts"`
	Database   string `json:"database"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Collection string `json:"collection"`
}

func ReadConfig() {
	file, _ := ioutil.ReadFile("model/../config.json")

	data := Config{}

	_ = json.Unmarshal([]byte(file), &data)
	Cfg = &data
}
