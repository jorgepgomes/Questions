package model

type Questions struct {
	_Id     interface{} `bson:_id`
	Id      int         `bson:"id"`
	Text    string      `bson:"text"`
	User    string      `bson:"username"`
	Date    int64       `bson:"creationDate"`
	Answers []Answers   `bson:"answers"`
}

type Answers struct {
	Id   int    `bson:"id"`
	Text string `bson:"text"`
	User string `bson:"user"`
	Date int64  `bson:"creationDate"`
}
