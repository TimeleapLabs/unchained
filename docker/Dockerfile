ARG NODE_VERSION
ARG UNCHAINED_VERSION
FROM node:${NODE_VERSION}
LABEL maintainer="hi@kenshi.io"

WORKDIR /app

RUN npm i -g @kenshi.io/unchained@${UNCHAINED_VERSION} && \
    unchained --version

ENTRYPOINT ["unchained", "start", "conf.yaml", "--generate"]
