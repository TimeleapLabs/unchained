#!/bin/sh

docker buildx build --build-arg="UBUNTU_VERSION=latest" --build-arg="UNCHAINED_VERSION=v0.11.0" -t ghcr.io/timeleaplabs/unchained:latest --no-cache .
