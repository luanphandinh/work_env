#!/usr/bin/env bash

usage() {
  echo "
usage: ${0##*/} [command]
commands:
  tmp     Clean up all temporary states(including temp env, profiles)
"
  exit 1
}

while [ "$1" != "" ]; do
  case $1 in
  tmp)
    shift
    > "${__TMP_DIR__}/.env"
    exit
    ;;

  -h | --help)
    usage
    exit
    ;;

  *)
    usage
    exit 1
    ;;
  esac
  shift
done
