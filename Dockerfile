FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

EXPOSE 5000

ENTRYPOINT [ "air" ]