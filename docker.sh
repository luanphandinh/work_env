#!/usr/bin/env bash

NAMES=$@
compose="docker-compose"

for NAME in $NAMES ; do
    compose+=" -f docker/compose/$NAME.yaml"
done
compose+=" up"

echo $compose
$compose
