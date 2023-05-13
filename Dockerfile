FROM --platform=${TARGETPLATFORM:-linux/amd64} node:18.16.0-alpine3.17 AS base

ARG TARGETPLATFORM

WORKDIR /app
COPY . .

RUN npm install -g pnpm
RUN pnpm install && pnpm build

FROM --platform=${TARGETPLATFORM:-linux/amd64} caddy:latest AS builder

COPY ./Caddyfile /etc/caddy/Caddyfile
COPY --from=base /app/build /srv

EXPOSE 8080

CMD ["caddy", "run", "--config=/etc/caddy/Caddyfile"]