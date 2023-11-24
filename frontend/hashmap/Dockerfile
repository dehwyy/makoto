FROM oven/bun as install
WORKDIR /src

COPY --from=ts/libs /app/libs libs
COPY --from=ts/libs /app/node_modules node_modules
COPY package.json .
COPY frontend/hashmap/package.json frontend/hashmap/package.json
RUN bun install

COPY frontend/hashmap frontend/hashmap/

FROM node:18-alpine
WORKDIR /app

COPY --from=install /src/libs libs
COPY --from=install /src/node_modules node_modules
COPY package.json .
COPY --from=install /src/frontend/hashmap frontend/hashmap/

RUN cd frontend/hashmap && npm run build

WORKDIR /app/frontend/hashmap/build

ENV PORT 3000
CMD [ "node", "." ]
