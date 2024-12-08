FROM golang:1.23.4 AS build

WORKDIR /date-service

COPY . .

RUN go build -o main .

EXPOSE 50003

CMD ["/date-service/main"]
