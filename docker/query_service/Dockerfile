FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .env ./
COPY ./cmd/query_service ./cmd
COPY ./internal/config ./internal/config
COPY ./internal/logger ./internal/logger
COPY ./internal/eventbus ./internal/eventbus
COPY ./internal/query_service ./internal/query_service
COPY ./internal/repository ./internal/repository
COPY ./internal/schema ./internal/schema
COPY ./internal/search ./internal/search
COPY ./internal/utils ./internal/utils

RUN go build -o ./bin/query_service ./cmd/main.go
EXPOSE 8080

ENTRYPOINT [ "bin/query_service" ]