FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o powertrain ./

FROM alpine:latest

COPY --from=builder /app/powertrain .

EXPOSE 9091

ENV APP_MODE=debug
ENV BROKER_PWD=mypassword
ENV BROKER_USR=myuser
ENV BROKER=rabbit

CMD ["./powertrain"]
