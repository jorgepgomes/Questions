import Request from "./request";
import config from "../config.json";

class Server {
  constructor() {
    this.request = new Request(config.client.url.server);
  }

  createQuestion = (data) => {
    return this.request.post("/api/questions/create", data);
  }
  AnswerQuestion = (data) => {
    return this.request.post("/api/questions/answers", data)
  }
  ListQuestions = () => {
    return this.request.get("/api/questions")
  }
  Like = (data) => {
    return this.request.post("/api/react", data)
  }
}

export default new Server();