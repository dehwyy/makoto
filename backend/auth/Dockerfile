FROM golang

WORKDIR /src

# caching dependencies
COPY ./grpc/go.mod ./grpc/go.sum ./grpc/
RUN cd grpc && go mod download
COPY ./auth/go.mod ./auth/go.sum ./auth/
RUN cd auth && go mod download


COPY ./grpc ./grpc
COPY ./_config /_config
COPY ./auth ./auth

RUN cd auth/cmd && go build -o main .

CMD ["./auth/cmd/main"]
