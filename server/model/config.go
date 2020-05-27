package model

type ConfigServer struct {
	ConfigMongoDB ConfigMongoDB
}

type ConfigMongoDB struct {
	Hosts      string `json:"hosts"`
	Database   string `json:"database"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Collection string `json:"collection"`
}
