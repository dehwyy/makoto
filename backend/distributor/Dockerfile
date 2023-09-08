FROM golang

WORKDIR /src

# caching
COPY ./distributor/go.mod ./distributor/go.sum ./distributor/

COPY ./distributor ./distributor
COPY ./grpc ./grpc
COPY ./_config /_config

RUN cd distributor/cmd && go build -o main .

CMD ["./distributor/cmd/main"]
