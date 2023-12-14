#!/bin/bash

if [ ! -d 'data' ]; then
    mkdir data
fi

UID_GID="$(id -u):$(id -g)" docker compose up -d
