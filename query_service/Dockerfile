FROM golang:1.20

ENV GOPATH /github.com
COPY . /github.com/salesforceanton/meower/query_service
WORKDIR /github.com/salesforceanton/meower/query_service

RUN go mod download
RUN go build -o ./bin/query_service ./cmd/query_service/main.go

ENTRYPOINT ["bin/query_service"]