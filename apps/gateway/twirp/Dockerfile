
FROM oven/bun as bun

RUN mkdir -p /temp/dev
COPY package.json /temp/dev
RUN mkdir -p /temp/dev/frontend
COPY libs /temp/dev/libs
RUN cd /temp/dev && bun install


FROM node:18-alpine as builder
RUN apk add build-base curl libc6-compat g++ git protoc

RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
COPY --from=golang:1.21.0-alpine /usr/local/go/ /usr/local/go/
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

WORKDIR /src

ENV PATH="$HOME/.cargo/bin:${PATH}"
ENV PATH="/root/.cargo/bin:${PATH}"

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1  && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
    go install github.com/twitchtv/twirp/protoc-gen-twirp@latest && \
    npm i -g pnpm

# copy files
COPY --from=bun /temp/dev/node_modules ./node_modules
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
