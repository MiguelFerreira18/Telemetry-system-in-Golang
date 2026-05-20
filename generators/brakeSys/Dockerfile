FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o brakes ./

FROM alpine:latest

COPY --from=builder /app/brakes .

EXPOSE 9094

ENV APP_MODE=debug
ENV BROKER_PWD=mypassword
ENV BROKER_USR=myuser
ENV BROKER=rabbit

CMD ["./brakes"]
