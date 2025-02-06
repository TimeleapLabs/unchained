ARG UBUNTU_VERSION
FROM ubuntu:${UBUNTU_VERSION}
LABEL maintainer="hi@timeleap.swiss"
ARG TARGETOS
ARG TARGETARCH
ARG TIMELEAP_VERSION

ARG DEBIAN_FRONTEND=noninteractive

WORKDIR /app

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install --no-install-recommends -y \
    curl ca-certificates && \
    curl -sSL https://github.com/TimeleapLabs/timeleap/releases/download/${TIMELEAP_VERSION}/timeleap.${TARGETOS}.${TARGETARCH} -o timeleap && \
    chmod +x timeleap && \
    apt-get remove -y --purge curl && \
    apt-get autoremove -y --purge && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY ./entrypoint.sh entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
