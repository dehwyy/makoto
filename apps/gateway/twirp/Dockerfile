FROM golang:1.21.2-alpine

WORKDIR /app

COPY --from=go/libs /app/libs libs
COPY apps/gateway/twirp/go* apps/gateway/twirp/
COPY apps/gateway/services/go* apps/gateway/services/

RUN go work init ./libs/config ./libs/logger ./libs/grpc ./apps/gateway/services ./apps/gateway/twirp

WORKDIR /app/apps/gateway/services
RUN go mod download
WORKDIR /app/apps/gateway/twirp
RUN go mod download

WORKDIR /app
COPY apps/gateway/services apps/gateway/services/
COPY apps/gateway/twirp apps/gateway/twirp/

RUN  go build -o gateway_twirp ./apps/gateway/twirp
CMD ["./gateway_twirp"]
