FROM golang:1.22.4 AS build

WORKDIR /profiles-service

COPY . .

RUN go build -o main .

CMD ["/profiles-service/main"]
