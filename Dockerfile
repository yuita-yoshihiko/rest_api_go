FROM golang:1.24.0-alpine

WORKDIR /rest-api-go
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git alpine-sdk
RUN set -x \
    && go mod download \
    && go install github.com/rubenv/sql-migrate/...@latest

COPY . .
