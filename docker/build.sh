#!/bin/sh

docker buildx build --build-arg="UBUNTU_VERSION=latest" -t ghcr.io/kenshitech/unchained:latest --no-cache .
