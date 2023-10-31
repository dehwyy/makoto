FROM golang:latest

WORKDIR /src

# context is a root of workspace

# workspace deps

# current
COPY ./apps/gateway/twirp/go.mod ./apps/gateway/twirp/go.sum ./apps/gateway/twirp/
COPY ./apps/gateway/twirp ./apps/gateway/twirp

# deps
COPY ./apps/gateway/services ./apps/gateway/services
COPY libs libs

RUN go work init ./apps/gateway/twirp ./apps/gateway/services ./libs ./libs/config ./libs/logger ./libs/database ./libs/grpc ./libs/middleware

RUN  go build -o gateway_twirp ./apps/gateway/twirp

CMD ["./gateway_twirp"]
