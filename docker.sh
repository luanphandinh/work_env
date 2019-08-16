#!/usr/bin/env bash

compose() {
	compose="docker-compose"
	SERVICES=$@

	for SERVICE in $SERVICES ; do
		compose+=" -f docker/compose/$SERVICE"
		if [[ $SERVICE != *".yaml"* ]]; then
			compose+=".yaml"
		fi
	done

	compose+=" up"
	echo $compose
}

NAMES=$@
COMMAND=""
if [ "$NAMES" == "all" ]; then
	COMMAND="$(compose `ls docker/compose`)"
else
	COMMAND="$(compose $NAMES)"
fi

echo $COMMAND
$COMMAND
