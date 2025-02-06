#!/bin/sh

echo "Running a $TIMELEAP_NODE_TYPE node."

./timeleap $TIMELEAP_CMD -c conf/conf.$TIMELEAP_NODE_TYPE.yaml -s conf/secrets.$TIMELEAP_NODE_TYPE.yaml -x context/$TIMELEAP_NODE_TYPE
