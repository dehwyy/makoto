FROM node:18-alpine as builder
COPY --from=golang:1.21.0-alpine /usr/local/go/ /usr/local/go/
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

WORKDIR /app

COPY go.work go.work.sum package.json tsconfig.json cargo.toml cargo.lock ./
COPY libs libs

# Backend microservices

RUN cd apps/auth \
  go mod download \
  go build -o ./out ./cmd/main.go
