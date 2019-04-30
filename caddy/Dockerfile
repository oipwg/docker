# Build image
FROM golang:1.12-alpine as builder

RUN apk add --no-cache git

ENV GO111MODULE=on

# import plugins
COPY with_plugins.go /build_caddy/with_plugins.go

WORKDIR /build_caddy
RUN go mod init caddy && \
    go get github.com/mholt/caddy && \
    go build

# Dist image
FROM alpine:3.6

# install deps
RUN apk add --no-cache --no-progress curl tini ca-certificates

# copy caddy binary
COPY --from=builder /build_caddy/caddy /usr/bin/caddy

# list plugins
RUN /usr/bin/caddy -plugins

# static files volume
VOLUME ["/www"]
WORKDIR /www

COPY Caddyfile /etc/caddy/Caddyfile

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["caddy", "-agree", "--conf", "/etc/caddy/Caddyfile"]