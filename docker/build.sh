#!/bin/sh

docker buildx build --build-arg="ALPINE_VERSION=latest" -t ghcr.io/kenshitech/unchained:latest --no-cache .
