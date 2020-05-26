#init server
FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3050

CMD ["./main"]