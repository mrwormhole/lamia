FROM --platform=${BUILDPLATFORM:-linux/amd64} node:alpine3.17 AS build

ARG BUILDPLATFORM
ARG TARGETPLATFORM

WORKDIR /app
COPY . .

RUN npm install -g pnpm
RUN pnpm install && pnpm build

FROM --platform=${TARGETPLATFORM:-linux/amd64} caddy:latest AS ship

COPY ./Caddyfile /etc/caddy/Caddyfile
COPY --from=build /app/build /srv

CMD ["caddy", "run", "--config=/etc/caddy/Caddyfile"]