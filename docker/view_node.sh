#!/bin/bash

UID_GID="$(id -u):$(id -g)" docker compose logs -f
