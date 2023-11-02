FROM golang:latest

WORKDIR /src

# context is a root of workspace

# current
COPY ./apps/auth/go.mod ./apps/auth/go.sum ./apps/auth/
COPY ./apps/auth ./apps/auth

# deps
COPY libs libs

RUN go work init ./apps/auth ./libs/config ./libs/logger ./libs/database ./libs/grpc

RUN  go build -o authorization ./apps/auth/cmd

CMD ["./authorization"]
