FROM golang:alpine as builder

ARG VERSION=dev

WORKDIR /backend

COPY ./ ./

RUN apk add --no-cache bash curl gcc git musl-dev 

RUN go build -o build/backend -ldflags "-s -w" -tags "jsoniter" -trimpath .

FROM alpine:latest

ENV PUID=0 PGID=0 UMASK=022

COPY --from=builder /backend/build/backend /usr/local/bin/backend

COPY script/entrypoint.sh /entrypoint.sh

RUN apk add --no-cache bash ca-certificates su-exec tzdata && \
    rm -rf /var/cache/apk/* && \
    chmod +x /entrypoint.sh && \
    mkdir -p /app

WORKDIR /app

EXPOSE 9000

ENTRYPOINT [ "/entrypoint.sh" ]

CMD []