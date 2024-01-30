ARG ALPINE_VERSION
FROM alpine:${ALPINE_VERSION}
LABEL maintainer="hi@kenshi.io"
ARG TARGETOS
ARG TARGETARCH
ARG UNCHAINED_VERSION

WORKDIR /app

RUN apk add --update curl && \
    curl -sSL https://github.com/KenshiTech/unchained/releases/download/go/unchained.${TARGETOS}.${TARGETARCH} -o unchained && \
    chmod +x unchained && \
    apk del curl && \
    rm -rf /var/cache/apk/*

COPY ./entrypoint.sh entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
