
FROM ubuntu as installer


# install dependencies
RUN apt-get -y update
RUN apt-get install -y curl protobuf-compiler unzip
RUN apt-get install -y --no-install-recommends build-essential
RUN apt-get update

RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y nodejs \
    npm

RUN npm install -g bun

COPY package.json ./
COPY frontend frontend
COPY libs libs

RUN bun install

COPY --from=golang:1.21.0-alpine /usr/local/go/ /usr/local/go/
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

WORKDIR /src


ENV PATH="$HOME/.cargo/bin:${PATH}"
ENV PATH="/root/.cargo/bin:${PATH}"

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1  && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
    go install github.com/twitchtv/twirp/protoc-gen-twirp@latest

# copy files
COPY ./apps/gateway/twirp/go.mod ./apps/gateway/twirp/go.sum ./apps/gateway/twirp/
COPY ./apps/gateway/twirp ./apps/gateway/twirp
COPY Cargo.toml .

# deps
COPY ./apps/gateway/services ./apps/gateway/services
COPY libs libs

RUN cd libs/grpc && cargo run

RUN go work init ./apps/gateway/twirp ./apps/gateway/services ./libs/config ./libs/logger ./libs/grpc


RUN  go build -o gateway_twirp ./apps/gateway/twirp

CMD ["./gateway_twirp"]
