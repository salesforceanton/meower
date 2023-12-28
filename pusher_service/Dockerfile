FROM golang:1.20

ENV GOPATH /github.com
COPY . /github.com/salesforceanton/meower/pusher_service
WORKDIR /github.com/salesforceanton/meower/pusher_service

RUN go mod download
RUN go build -o ./bin/pusher_service ./cmd/pusher_service/main.go

ENTRYPOINT ["bin/pusher_service"]