FROM golang:1.20

ENV GOPATH /github.com
COPY . /github.com/salesforceanton/meower/meow_service
WORKDIR /github.com/salesforceanton/meower/meow_service

RUN go mod download
RUN go build -o ./bin/meow_service ./cmd/meow_service/main.go

CMD ["./bin/meow_service"]