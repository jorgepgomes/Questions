package model

type Questions struct {
	_Id     interface{} `bson:_id`
	Id      int         `bson:"id"`
	Text    string      `bson:"text"`
	User    string      `bson:"username"`
	Likes   int         `bson:"likes"`
	Date    int64       `bson:"creationDate"`
	Answers []Answers   `bson:"answers"`
}

type Answers struct {
	Id    int    `bson:"id"`
	Text  string `bson:"text"`
	User  string `bson:"user"`
	Likes int    `bson:"likes"`
	Date  int64  `bson:"creationDate"`
}

type Like struct {
	IdQuestion int    `json:"id_question"`
	IdAnswer   int    `json:"id_answer"`
	Local      string `json:"local"`
	Like       int    `json:"like"`
}
