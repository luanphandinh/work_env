#!/usr/bin/env bash

info() {
  echo ""
  echo "Usage: $0 -e ENV -s PARAMS"
  echo -e "\t-e e Environment: Possible values ['dev', 'test'], avoiding conflict data, accidentally delete dev data when running test"
  echo -e "\t-s s Services: optional services to bootstrap"
  exit 1
}

compose() {
  compose="docker-compose"
  PARAMS=$@

  for SERVICE in ${PARAMS}; do
    compose+=" -f docker/compose/$SERVICE"
    if [[ ${SERVICE} != *".yaml"* ]]; then
      compose+=".yaml"
    fi
  done

  compose+=" up"
  echo "$compose"
}

command() {
  PARAMS=$@
  SERVICES="all"

  if [[ -n "$PARAMS" ]]; then
    FILTER=""
    for SERVICE in ${PARAMS}; do
      FILTER+="-e $SERVICE.*yaml "
    done

    SERVICES=$(ls -p docker/compose | grep ${FILTER})
  fi

  COMMAND=""
  if [[ "$SERVICES" == "all" ]]; then
    COMMAND="$(compose $(ls docker/compose))"
  else
    COMMAND="$(compose "$SERVICES")"
  fi

  echo "$COMMAND"
  ${COMMAND}
}

env() {
  ENV=$1

  if [[ -z "$ENV" ]]; then
    ENV="dev"
  fi

  if [[ "$ENV" != "dev" ]] && [[ "$ENV" != "test" ]]; then
    echo "-e only accept values ['dev', 'test']"
    exit 1
  fi

  echo "Docker containers running on ENV: $ENV"
  export ENV="$ENV"
}

while getopts "e:s:h:" opt; do
  case "$opt" in
  e) ENV="$OPTARG" ;;
  s) PARAMS="$OPTARG" ;;
  h) info ;;
  esac
done

env "$ENV"
command "$PARAMS"
