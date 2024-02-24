#!/bin/sh

echo "Running a $UNCHAINED_NODE_TYPE node."

./unchained $UNCHAINED_CMD -c conf/conf.$UNCHAINED_NODE_TYPE.yaml -s conf/secrets.$UNCHAINED_NODE_TYPE.yaml -x context/$UNCHAINED_NODE_TYPE
