FROM golang:1.22.4 AS build

WORKDIR /users-service

COPY . .

RUN go build -o main .

CMD ["/users-service/main"]
