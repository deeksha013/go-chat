FROM golang:latest


LABEL maintainer="dee <hello@world>"

RUN mkdir /app

WORKDIR /app


RUN go mod download

COPY . .

ENV PORT ":8080"

RUN go build -o main

CMD ["/app/main"]
