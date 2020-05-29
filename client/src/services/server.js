import Request from "./request";
import config from "../config.json";

class Server {
  constructor() {
    this.request = new Request(config.client.url.server);
  }

  CreateQuestion = (data) => {
    return this.request.post("/api/questions/create", data);
  }
  AnswerQuestion = (id, data) => {
    return this.request.post(`/api/questions/answers?id=${id}`, data)
  }
  ListQuestions = (data) => {
    return this.request.get(`/api/questions?search=${data}`)
  }
  Like = (data) => {
    return this.request.post("/api/react", data)
  }

  DetailsQuestion = (data) => {
    return this.request.get(`/api/questions/details?id=${data}`)
  }
}

export default new Server();