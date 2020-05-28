package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/jorgepgomes/Questions/server/model"
	"gopkg.in/mgo.v2/bson"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	// setupResponse(&w, r)
	body := getBody(r)

	count := totalQuestions()

	id := generateID(count)

	var question model.Questions
	_ = json.Unmarshal(body, &question)

	question.Id = id
	question.Date = dateNow()
	question.Likes = 0

	response := insertQuestion(question)

	w.WriteHeader(response.Code)
	w.Write([]byte(model.ToJson(response)))
}

func AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	// setupResponse(&w, r)
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
	answer.Likes = 0

	response := pushAnswer(id, answer)

	w.WriteHeader(response.Code)
	w.Write([]byte(model.ToJson(response)))
}

func ListQuestions(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	search := getSearchParam(r)

	result := findQuestions(search)

	w.WriteHeader(200)
	w.Write([]byte(model.ToJson(result)))
}

func Like(w http.ResponseWriter, r *http.Request) {
	// setupResponse(&w, r)
	body := getBody(r)

	var like model.Like
	err := json.Unmarshal(body, &like)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	question := findOneQuestion(like.IdQuestion)

	response := updateLike(question, like)

	w.WriteHeader(200)
	w.Write([]byte(model.ToJson(response)))
}

func findQuestions(search string) []model.Questions {
	col := mongoStore.session.DB(model.Cfg.Server.MongoDB.Database).C(model.Cfg.Server.MongoDB.Collection)
	results := []model.Questions{}
	err := col.Find(bson.M{"text": bson.RegEx{search, ""}}).All(&results)
	if err != nil {
		return nil
	}
	return results
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

func updateLike(questions *model.Questions, likes model.Like) *model.Response {
	response := messageResponse(200, "Like salvo com sucesso")
	col := mongoStore.session.DB(model.Cfg.Server.MongoDB.Database).C(model.Cfg.Server.MongoDB.Collection)
	var err error
	total := 0
	if likes.Local == "question" {
		like := questions.Likes
		total = like + likes.Like
		if total < 0 {
			total = 0
		}
	} else {
		positionArray := likes.IdAnswer - 1
		like := questions.Answers[positionArray].Likes
		total = like + likes.Like
		if total < 0 {
			total = 0
		}
	}

	if likes.Local == "question" {
		err = col.Update(bson.M{"id": likes.IdQuestion}, bson.M{"$set": bson.M{"likes": total}})
	} else {
		err = col.Update(bson.M{"id": likes.IdQuestion, "answers.id": likes.IdAnswer}, bson.M{"$set": bson.M{"answers.$.likes": total}})
	}
	if err != nil {
		fmt.Println(err)
		response = messageResponse(400, "Ocorreu um erro ao salvar o like")
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

func getSearchParam(r *http.Request) string {
	param := r.URL.Query()["search"]
	if len(param) > 0 {
		return param[0]
	}
	return ""
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
