FROM oven/bun as install
WORKDIR /src

RUN mkdir -p /temp/prod
COPY package.json tsconfig.json /temp/prod/
COPY ./frontend/auth /temp/prod/frontend/auth
COPY Cargo.toml /temp/prod/
COPY libs/config /temp/prod/libs/config
COPY libs/grpc /temp/prod/libs/grpc
COPY libs/lib /temp/prod/libs/lib

RUN cd /temp/prod && bun install
RUN cd /temp/prod/libs/config && bun run ts
RUN cd /temp/prod/libs/lib && bun run ts


FROM node:18-alpine as builder
WORKDIR /src

RUN apk add build-base curl libc6-compat g++ git protoc
# cargo install
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
ENV PATH="$HOME/.cargo/bin:${PATH}"
ENV PATH="/root/.cargo/bin:${PATH}"

# go
COPY --from=golang:1.21.0-alpine /usr/local/go/ /usr/local/go/
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

# install deps for gprc(twirp gen)
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1  && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
    go install github.com/twitchtv/twirp/protoc-gen-twirp@latest

COPY --from=install /temp/prod/node_modules node_modules
COPY --from=install /temp/prod/libs libs
COPY package.json .
COPY tsconfig.json .
COPY ./frontend/auth ./frontend/auth
COPY Cargo.toml .


# generate files go/ts
RUN cd libs/grpc && cargo run

RUN cd frontend/auth && npm run build

WORKDIR /src/frontend/auth/build


ENV PORT 3001
CMD [ "node", "." ]
