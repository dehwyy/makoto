FROM oven/bun as bun

# node_modules
RUN mkdir -p /temp/dev
COPY package.json /temp/dev
RUN mkdir -p /temp/dev/frontend
COPY libs /temp/dev/libs
RUN cd /temp/dev && bun install

# app build
FROM node:18-alpine as builder

# deps
RUN apk add build-base curl libc6-compat g++ git protoc

# cargo install
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
ENV PATH="$HOME/.cargo/bin:${PATH}"
ENV PATH="/root/.cargo/bin:${PATH}"

# go
COPY --from=golang:1.21.0-alpine /usr/local/go/ /usr/local/go/
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

#
WORKDIR /src

# install deps for gprc(twirp gen)
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1  && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
    go install github.com/twitchtv/twirp/protoc-gen-twirp@latest && \
    npm i -g pnpm

# copy files
COPY --from=bun /temp/dev/node_modules ./node_modules
COPY ./apps/auth/go.mod ./apps/auth/go.sum ./apps/auth/
COPY ./apps/auth ./apps/auth
COPY Cargo.toml .
COPY libs libs

# generate files go/ts
RUN cd libs/grpc && cargo run

# create go workspace only with necessary packages
RUN go work init ./apps/auth ./libs/database ./libs/config ./libs/logger ./libs/grpc


RUN  go build -o auth_backend ./apps/auth/cmd

CMD ["./auth_backend"]
