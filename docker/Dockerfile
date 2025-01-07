ARG UBUNTU_VERSION
FROM ubuntu:${UBUNTU_VERSION}
LABEL maintainer="hi@timeleap.swiss"
ARG TARGETOS
ARG TARGETARCH
ARG UNCHAINED_VERSION

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /app

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install --no-install-recommends -y \
    curl ca-certificates && \
    curl -sSL https://github.com/TimeleapLabs/unchained/releases/download/${UNCHAINED_VERSION}/unchained.${TARGETOS}.${TARGETARCH} -o unchained && \
    chmod +x unchained && \
    apt-get remove -y --purge curl && \
    apt-get autoremove -y --purge && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY ./entrypoint.sh entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
