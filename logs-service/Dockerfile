FROM golang:1.23.4 AS build

WORKDIR /logs-service

COPY . .

RUN go build -o main .

EXPOSE 50002

CMD ["/logs-service/main"]
