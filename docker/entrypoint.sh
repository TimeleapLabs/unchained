#!/bin/sh

echo "Running a $UNCHAINED_NODE_TYPE node."

#if [ $UNCHAINED_NODE_TYPE = "broker" ]; then
#  unchained postgres migrate conf.yaml
#  retVal=$?
#  if [ $retVal -ne 0 ]; then
#    exit $retVal
#  fi
#fi

./unchained $UNCHAINED_NODE_TYPE -c conf/conf.$UNCHAINED_NODE_TYPE.yaml -s conf/secrets.$UNCHAINED_NODE_TYPE.yaml
