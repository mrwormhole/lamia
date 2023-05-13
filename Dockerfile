from node:18.16.0-alpine3.17 AS base

WORKDIR /app
COPY . .

RUN npm install -g pnpm
RUN pnpm install && pnpm build

from caddy:latest AS builder

COPY ./Caddyfile /etc/caddy/Caddyfile
COPY --from=base /app/build /srv

EXPOSE 8080

CMD ["caddy", "run", "--config=/etc/caddy/Caddyfile"]