FROM golang:1.17 AS builder

RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /github.com/salesforceanton/meower


COPY config config
COPY logger logger
COPY utils utils
COPY schema schema
COPY eventbus eventbus
COPY repository repository
COPY search search
COPY web_socket web_socket
COPY meow_service meow_service
COPY query_service query_service
COPY pusher_service pusher_service

RUN go install ./...

FROM alpine:latest
WORKDIR /usr/bin
COPY --from=build /go/bin .