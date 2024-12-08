FROM golang:1.23.4 AS build

WORKDIR /payment-service

COPY . .

RUN go build -o main .

EXPOSE 50005

CMD ["/payment-service/main"]
