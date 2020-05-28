package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/jorgepgomes/Questions/server/model"
	"gopkg.in/mgo.v2/bson"
)

func CreateQuestion(w http.ResponseWriter, r *http.Request) {

	body := getBody(r)

	count := totalQuestions()

	id := generateID(count)

	var question model.Questions
	_ = json.Unmarshal(body, &question)

	question.Id = id
	question.Date = dateNow()

	response := insertQuestion(question)

	w.WriteHeader(response.Code)
	w.Write([]byte(model.ToJson(response)))
}

func AnswerQuestion(w http.ResponseWriter, r *http.Request) {

	body := getBody(r)
	id := getIDParams(r)
	question := findOneQuestion(id)
	idAnswer := generateID(len(question.Answers))

	var answer model.Answers
	err := json.Unmarshal(body, &answer)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	answer.Id = idAnswer
	answer.Date = dateNow()

	response := pushAnswer(id, answer)

	w.WriteHeader(response.Code)
	w.Write([]byte(model.ToJson(response)))
}

func findOneQuestion(id int) *model.Questions {
	col := mongoStore.session.DB(model.Cfg.Server.MongoDB.Database).C(model.Cfg.Server.MongoDB.Collection)
	var result model.Questions
	err := col.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		return nil
	}
	return &result
}

func pushAnswer(id int, answer model.Answers) *model.Response {
	response := messageResponse(200, "Pergunta respondida com sucesso")
	col := mongoStore.session.DB(model.Cfg.Server.MongoDB.Database).C(model.Cfg.Server.MongoDB.Collection)
	change := bson.M{"$push": bson.M{"answers": answer}}
	err := col.Update(bson.M{"id": id}, change)
	if err != nil {
		response = messageResponse(400, "Ocorreu um erro ao salvar sua resposta")
	}
	return response
}

func insertQuestion(question model.Questions) *model.Response {
	response := messageResponse(200, "Pergunta criada com sucesso")
	col := mongoStore.session.DB(model.Cfg.Server.MongoDB.Database).C(model.Cfg.Server.MongoDB.Collection)
	err := col.Insert(question)
	if err != nil {
		response = messageResponse(200, "Ocorreu um erro ao criar sua pergunta")
	}
	return response
}

func generateID(position int) int {
	if position > 0 {
		position = position + 1
	} else {
		position = 1
	}
	return position
}

func getBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil
	}
	return body
}

func getIDParams(r *http.Request) int {
	questionID := r.URL.Query()["id"]
	id, _ := strconv.Atoi(questionID[0])
	return id
}

func dateNow() int64 {
	dateNow := time.Now().Unix() * 1000
	return dateNow
}

func messageResponse(status int, message string) *model.Response {
	response := &model.Response{
		Code:    status,
		Message: message,
	}
	return response
}

func totalQuestions() int {
	col := mongoStore.session.DB(model.Cfg.Server.MongoDB.Database).C(model.Cfg.Server.MongoDB.Collection)
	count, err := col.Count()
	if err != nil {
		panic(err)
	}
	return count
}
