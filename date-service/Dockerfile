FROM golang:1.22.4 AS build

WORKDIR /date-service

COPY . .

RUN go build -o main .

CMD ["/date-service/main"]
