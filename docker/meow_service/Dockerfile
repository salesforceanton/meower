FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .env ./
COPY ./cmd/meow_service ./cmd
COPY ./internal/config ./internal/config
COPY ./internal/logger ./internal/logger
COPY ./internal/eventbus ./internal/eventbus
COPY ./internal/meow_service ./internal/meow_service
COPY ./internal/repository ./internal/repository
COPY ./internal/schema ./internal/schema
COPY ./internal/utils ./internal/utils

RUN go build -o ./bin/meow_service ./cmd/main.go
EXPOSE 8080

ENTRYPOINT [ "bin/meow_service" ]