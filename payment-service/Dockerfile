FROM golang:1.22.4 AS build

WORKDIR /payment-service

COPY . .

RUN go build -o main .

CMD ["/payment-service/main"]
