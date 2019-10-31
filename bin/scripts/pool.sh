#!/usr/bin/env bash

TIMES=$1
REQUEST=$2

for i in `seq 0 $TIMES`
do
	echo "call $i times"
	curl $REQUEST &
done
