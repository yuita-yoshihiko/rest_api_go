FROM golang:1.24.0-alpine

WORKDIR /rest-api-go
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git alpine-sdk
RUN set -x \
    && go mod download \
    && go install github.com/rubenv/sql-migrate/...@latest \
    && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest \
    && go install github.com/air-verse/air@latest

COPY . .
