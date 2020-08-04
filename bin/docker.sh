#!/usr/bin/env bash

declare -xr __DOCKER_PATH__="$__VAR_LIB_DIR__/docker/$__PROFILE__"
declare -xr __DOCKER_SERVICE_PREFIX__="cli_${__PROFILE__}"

usage() {
  echo "
usage: ${0##*/} { run name... | stop | clean [ --cache ] }
commands:

    run     Start all specify docker containers.
            eg: docker run mysql adminer

            Options:

            [--like]    To run services that contain keyword
                        eg: docker run --like sql

    stop    Stop all docker containers within profile's name.(<default> is used if not specify).

    clean   Clean up all docker containers within profile's name (<default> is used if not specify).

            Options:

            [--cache]       To remove cache.
"
  exit 1
}

#######################################
# Compose service name, finding all docker-compose that exist in __ENV_ROOT__/etc/docker and compose command
# Globals:
#   __ENV_ROOT__
#   __LOG__
#   DEBUG
# Arguments:
#   service_names{...}
#   ex: adminer mysql
# Returns:
#   docker-compose commmand
#   ex:  docker-compose -f /Users/phanluan/env/etc/docker/adminer.yaml -f /Users/phanluan/env/etc/docker/mysql.yaml up
#######################################
compose() {
  compose="docker-compose"
  PARAMS=$@

  for SERVICE in ${PARAMS}; do
    compose+=" -f ${__ENV_ROOT__}/etc/docker/${SERVICE}"
    if [[ ${SERVICE} != *".yaml"* ]]; then
      compose+=".yaml"
    fi
  done

  compose+=" up"
  if [[ $ENV_SILENT -eq 1 ]]; then
    compose+=" -d"
  fi

  echo "${compose}"
}

#######################################
# Up and running docker services
# Globals:
#   __ENV_ROOT__
#   __LOG__
# Arguments:
#   [--like] service_names{...}
#   ex: adminer mysql
# Returns:
#   docker-compose commmand
#   ex:  docker-compose -f /Users/phanluan/env/etc/docker/adminer.yaml -f /Users/phanluan/env/etc/docker/mysql.yaml up
#######################################
run() {
  if [[ "$1" == "all" ]]; then
    COMMAND=$(compose $(ls -p "${__ENV_ROOT__}/etc/docker" | grep -v /))
  else
    SERVICES=$@
    if [[ "$1" == "--like" ]]; then
      PARAMS=$@

      if [[ -n "${PARAMS}" ]]; then
        FILTER=""
        for SERVICE in ${PARAMS}; do
          FILTER+="-e ${SERVICE}.*yaml "
        done

        SERVICES=$(ls -p "${__ENV_ROOT__}/etc/docker" | grep "${FILTER}")
      fi
    fi
    COMMAND=$(compose "${SERVICES}")
  fi

  $__LOG__ -i "${COMMAND}"
  ${COMMAND}
}

#######################################
# Stop all runing container that contain strings
# Use case in this project is for stopping all docker image that name contains cli_<profile_name>.
# Globals:
#   __DOCKER_SERVICE_PREFIX__
# Arguments:
#   None
# Returns:
#   None
#######################################
stop() {
  $__LOG__ -i stopped: $(docker stop $(docker ps --format "{{.Names}}" | grep "${__DOCKER_SERVICE_PREFIX__}"))
}

#######################################
# Clean all container that contain strings
# Use case in this project is for cleaning all docker image that name contains cli_<profile_name>.
# Globals:
#   __LOG__
#   __DOCKER_SERVICE_PREFIX__
#   __DOCKER_PATH__
# Arguments:
#   None
# Returns:
#   None
#######################################
clean() {
  if [[ $(docker ps -a --format "{{.Names}}" | grep "${__DOCKER_SERVICE_PREFIX__}" | wc -l) -gt 0 ]]; then
    $__LOG__ -i "docker clean up images."
    $__LOG__ -i "cleaned: $(docker rm $(docker ps -a --format "{{.Names}}" | grep "${__DOCKER_SERVICE_PREFIX__}"))"
    $__LOG__ -i "docker clean up images done."
  fi

  if [ "$1" == "--cache" ]; then
    $__LOG__ -i "docker clean up images cache."
    rm -rf "${__DOCKER_PATH__}"
    $__LOG__ -i "docker clean up images cache done."
  fi
}

while [ "$1" != "" ]; do
  case $1 in
  run)
    shift
    if [[ ! -d "${__DOCKER_PATH__}" ]]; then
      mkdir "${__DOCKER_PATH__}"
    fi
    $__LOG__ -i "docker runs: $@"
    run $@
    exit
    ;;

  stop)
    shift
    $__LOG__ -i "docker stop."
    stop
    exit
    ;;

  clean)
    shift
    clean $1
    exit
    ;;

  list|ls)
    shift
    ls $__DOCKER_DIR__
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
