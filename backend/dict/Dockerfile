FROM golang:latest

# caching dependencies

# grpc
COPY ./grpc/go.mod ./grpc/go.sum ./grpc/
RUN cd grpc && go mod download

# current module
COPY ./dict/go.mod ./dict/go.sum ./dict/
RUN cd dict && go mod download


COPY ./grpc ./grpc
COPY ./_config /_config
COPY ./dict ./dict

RUN cd dict/cmd && go build -o main .

CMD ["./dict/cmd/main"]
