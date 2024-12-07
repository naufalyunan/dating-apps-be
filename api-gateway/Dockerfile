FROM golang:1.23.4 AS build

WORKDIR /api-gateway

COPY . .

RUN go build -o main .

CMD ["/api-gateway/main"]
