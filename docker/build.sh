#!/bin/sh

docker buildx build --build-arg="UBUNTU_VERSION=latest" --build-arg="TIMELEAP_VERSION=v0.11.0" -t ghcr.io/timeleaplabs/timeleap:latest --no-cache .
