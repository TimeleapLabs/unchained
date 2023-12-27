#!/bin/bash

usage() {
  echo "Usage: $0 [node] [options]"
  echo "Node:"
  echo "  local  - Manage unchained local node"
  echo "  remote - Manage unchained remote node"
  echo "  lite   - Manage unchained lite node"
  echo "Options:"
  echo "  Additional options passed directly to 'docker compose'"
  echo "Examples:"
  echo "  To start a node: $0 local up -d"
  echo "  To stop a node: $0 local stop"
  echo "  To view the status of a node: $0 local ps"
  echo "  To view logs of a node: $0 local logs -f"
}

if ! command -v docker &>/dev/null; then
  echo "Error: docker could not be found on your system!"
  exit 1
fi

if ! docker compose version &>/dev/null; then
  echo "Error: docker compose could not be found on your system!"
  exit 1
fi

if [ ! $1 == 'local' ] && [ ! $1 == 'remote' ] && [ ! $1 == 'lite' ] || [ -z $2 ]; then
  usage
  exit 1
fi

UID_GID="$(id -u):$(id -g)"

if [ $2 == 'up' ] && [ $1 == 'local' ] && [ ! -d 'data' ]; then
  mkdir data
fi

UID_GID=$UID_GID COMPOSE_PROFILES=$1 docker compose "${@:2}"
