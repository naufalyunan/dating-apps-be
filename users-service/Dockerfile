FROM golang:1.23.4 AS build

WORKDIR /users-service

COPY . .

RUN go build -o main .

EXPOSE 50001

CMD ["/users-service/main"]
