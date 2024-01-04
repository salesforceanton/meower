FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .env ./
COPY ./cmd/pusher_service ./cmd
COPY ./internal/config ./internal/config
COPY ./internal/logger ./internal/logger
COPY ./internal/eventbus ./internal/eventbus
COPY ./internal/pusher_service ./internal/pusher_service
COPY ./internal/schema ./internal/schema
COPY ./internal/web_socket ./internal/web_socket
COPY ./internal/utils ./internal/utils

RUN go build -o ./bin/pusher_service ./cmd/main.go
EXPOSE 8080

ENTRYPOINT [ "bin/pusher_service" ]