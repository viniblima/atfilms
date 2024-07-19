FROM golang:1.17

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy