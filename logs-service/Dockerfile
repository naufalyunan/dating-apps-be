FROM golang:1.22.4 AS build

WORKDIR /logs-service

COPY . .

RUN go build -o main .

CMD ["/logs-service/main"]
