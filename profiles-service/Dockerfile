FROM golang:1.23.4 AS build

WORKDIR /profiles-service

COPY . .

RUN go build -o main .

EXPOSE 50004

CMD ["/profiles-service/main"]
