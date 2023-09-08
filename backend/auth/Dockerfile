FROM golang

WORKDIR /src

# caching dependencies
COPY ./auth/go.mod ./auth/go.sum ./auth/
COPY ./grpc/go.mod ./grpc/go.sum ./grpc/

COPY ./auth ./auth
COPY ./grpc ./grpc
COPY ./_config /_config

RUN cd auth/cmd && go build -o main .

CMD ["./auth/cmd/main"]
