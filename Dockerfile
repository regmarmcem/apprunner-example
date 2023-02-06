# syntax=docker/dockerfile:1
FROM golang:1.19.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN apk update && go mod download

COPY . ./

RUN go build -o /taskapp

EXPOSE 8080

CMD ["./taskapp"]

