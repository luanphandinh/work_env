#!/usr/bin/env bash

# All available global variables
# Naming convention:
# Variables should be __<NAME>__
# Directory should be __<NAME>_DIR__
# Exec file should be __<NAME>_EXEC__
# Variables
declare -xr __ENV_ROOT__=$(cd $(dirname $0) && pwd)
declare -xr __CLI_VERSION__=$(cat $__ENV_ROOT__/VERSION)
declare -x __PROFILE__="default"

# Directory
declare -xr __DOCKER_DIR__="${__ENV_ROOT__}/etc/docker"
declare -xr __PROFILE_DIR__="${__ENV_ROOT__}/etc/profile"
declare -xr __VAR_LOG_DIR__="${__ENV_ROOT__}/var/log"
declare -xr __VAR_LIB_DIR__="${__ENV_ROOT__}/var/lib"
declare -xr __VAR_MAIL_DIR__="${__ENV_ROOT__}/var/mail"

# bin/**
declare -xr __DOCKER_EXEC__="${__ENV_ROOT__}/bin/docker.sh"
declare -xr __CONFIG_PROFILE_EXEC__="${__ENV_ROOT__}/bin/config_profile.sh"
declare -xr __LOG__="${__ENV_ROOT__}/bin/log.sh"
declare -xr __DEBUG__="${__ENV_ROOT__}/bin/debug.sh"

# Export environment ports from etc/docker/
set -a
. "${__DOCKER_DIR__}/.PORT"

usage() {
  cli_name=${0##*/}
  echo "
${cli_name}
your profile CLI
version: ${__CLI_VERSION__}
usage: ${cli_name} [options] [command [command's options]]

options:
  -p | --profile <profile_name>:  Profile that cli with take action on.
                                  Auto create new one if not exist.
                                  defualt <profile_name>: default.

  -d | --debug:                   Turn on debug mode.

  -h | --help:                    Help.

commands:
  docker:           Up and running dockers container
                    All possible containers a listed in etc/docker
                    Will create volume for corresponding containers in var/lib/<profile_name>/docker

  set:              Config profile.
                    eg: ./cli.sh -p luanphan set SOME_VAR=SOME_VALUE OTHER_VAR=OVER_VALUE

  checkconf:        printenv of current profile to screen.
"
  exit 1
}

#######################################
# Apply profile ENV variables to current shell and it's childs
# Globals:
#   __ENV_ROOT__
#   __PROFILE__
# Arguments:
#   None
# Returns:
#   None
#######################################
apply_profile_config() {
  __dir__="${__ENV_ROOT__}/etc/profile/${__PROFILE__}"
  if [[ ! -d "${__dir__}" ]]; then
    mkdir "${__dir__}"
  fi

  export __profile_config_file__="${__dir__}/.env"
  if [[ ! -d "${__profile_config_file__}" ]]; then
    touch "${__profile_config_file__}"
  fi

  $__DEBUG__ "EXPORT ENV from profile: ${__PROFILE__}"
  $__DEBUG__ $(cat ${__profile_config_file__})

  set -a
  . "${__profile_config_file__}"
}

while [ "$1" != "" ]; do
  case $1 in
  -p | --profile)
    shift
    set -e
    export __PROFILE__=$1
    ;;

  -d | --debug)
    set -e
    export DEBUG=1
    ;;

  run)
    shift
    apply_profile_config
    $@
    exit;;

  checkenv)
    apply_profile_config
    printenv
    exit;;

  checkconf)
    shift
    apply_profile_config
    $__CONFIG_PROFILE_EXEC__ -n $__PROFILE__ checkconf $@
    exit;;

  docker)
    shift
    $__LOG__ -i "CLI running on profile: ${__PROFILE__}"
    apply_profile_config
    $__DOCKER_EXEC__ $@
    exit;;

  set)
    shift
    $__CONFIG_PROFILE_EXEC__ -n $__PROFILE__ set $@
    exit;;

  config-profile)
    shift
    $__CONFIG_PROFILE_EXEC__ $@
    exit;;

  -h | --help)
    usage
    exit;;

  *)
    usage
    exit 1;;
  esac
  shift
done
