#!/bin/bash

usage() {
  echo "Usage: $0 [node] [options]"
  echo "Node:"
  echo "  local - Manage unchained local node"
  echo "  atlas - Manage unchained atlas node"
  echo "  lite  - Manage unchained lite node"
  echo "Options:"
  echo "  Additional options passed directly to 'docker compose'"
  echo "Examples:"
  echo "  To start a node: $0 local up -d"
  echo "  To stop a node: $0 local stop"
  echo "  To view the status of a node: $0 local ps"
  echo "  To view logs of a node: $0 local logs -f"
}

if [ ! $1 == 'local' ] && [ ! $1 == 'atlas' ] && [ ! $1 == 'lite' ] ; then
  usage
  exit 1
fi

UID_GID="$(id -u):$(id -g)"

if [ $2 == 'up' ] ; then 
  if [ $1 == 'local' ] && [ ! -d 'data' ]; then
    mkdir data
  fi
fi

UID_GID=$UID_GID COMPOSE_PROFILES=$1 docker compose "${@:2}"
