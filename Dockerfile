FROM --platform=${BUILDPLATFORM:-linux/amd64} node:alpine3.17 AS build

ARG BUILDPLATFORM
ARG TARGETPLATFORM

WORKDIR /app
COPY . .

RUN npm install -g pnpm
RUN pnpm install && pnpm build

FROM --platform=${TARGETPLATFORM:-linux/amd64} nginx:latest AS ship

COPY ./nginx.conf etc/nginx/conf.d/default.conf
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80

STOPSIGNAL SIGQUIT

CMD ["nginx", "-g", "daemon off;"]