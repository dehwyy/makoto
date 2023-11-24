FROM golang:1.21.2-alpine

WORKDIR /app

COPY --from=go/libs /app/libs libs
COPY apps/auth/go* apps/auth/

RUN go work init ./libs/config ./libs/logger ./libs/grpc ./libs/database ./apps/auth

WORKDIR /app/apps/auth
RUN go mod download

WORKDIR /app
COPY apps/auth apps/auth/

RUN  go build -o authorization ./apps/auth/cmd
CMD ["./authorization"]
