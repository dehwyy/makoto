FROM oven/bun as install
WORKDIR /src

COPY --from=ts/libs /app/libs libs
COPY --from=ts/libs /app/node_modules node_modules
COPY package.json .
COPY frontend/main/package.json frontend/main/package.json
RUN bun install

COPY frontend/main frontend/main/

FROM node:18-alpine
WORKDIR /app

COPY --from=install /src/libs libs
COPY --from=install /src/node_modules node_modules
COPY package.json .
COPY --from=install /src/frontend/main frontend/main/

RUN cd frontend/main && npm run build

WORKDIR /app/frontend/main

ENV NODE_ENV production
ENV NEXT_TELEMETRY_DISABLED 1
ENV PORT 3000

CMD [ "npm", "run", "start" ]
