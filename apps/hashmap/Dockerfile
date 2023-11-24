FROM golang:1.21.2-alpine

WORKDIR /app

COPY --from=go/libs /app/libs libs
COPY apps/hashmap/go* apps/hashmap/

RUN go work init ./libs/config ./libs/logger ./libs/grpc ./libs/database ./libs/lib ./apps/hashmap

WORKDIR /app/apps/hashmap
RUN go mod download

WORKDIR /app
COPY apps/hashmap apps/hashmap/

RUN  go build -o hashmap ./apps/hashmap/cmd
CMD ["./hashmap"]
