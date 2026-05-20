FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o climatecontrols ./

FROM alpine:latest

COPY --from=builder /app/climatecontrols .

EXPOSE 9093

ENV APP_MODE=debug
ENV BROKER_PWD=mypassword
ENV BROKER_USR=myuser
ENV BROKER=rabbit

CMD ["./climatecontrols"]
