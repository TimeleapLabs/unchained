#!/bin/sh

if [ "$UNCHAINED_NODE_TYPE" == "full" ]; then
  unchained postgres migrate conf.yaml
fi

unchained start conf.yaml --generate
