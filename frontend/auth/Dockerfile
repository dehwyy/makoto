FROM oven/bun as install
WORKDIR /src

COPY --from=ts/libs /app/libs libs
COPY --from=ts/libs /app/node_modules node_modules
COPY package.json .
COPY frontend/auth/package.json frontend/auth/package.json
RUN bun install

COPY frontend/auth frontend/auth/

FROM node:18-alpine
WORKDIR /app

COPY --from=install /src/libs libs
COPY --from=install /src/node_modules node_modules
COPY package.json .
COPY --from=install /src/frontend/auth frontend/auth/

RUN cd frontend/auth && npm run build

WORKDIR /app/frontend/auth/build

ENV PORT 3000
CMD [ "node", "." ]
