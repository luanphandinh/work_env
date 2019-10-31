#!/usr/bin/env bash

docker stop `docker ps --format "{{.Names}}" | grep "work"`
