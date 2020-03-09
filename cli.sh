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
declare -xr __TMP_DIR__="${__ENV_ROOT__}/tmp"

# bin/**
declare -xr __DOCKER_EXEC__="${__ENV_ROOT__}/bin/docker.sh"
declare -xr __RUNNER_EXEC__="${__ENV_ROOT__}/bin/runner.sh"
declare -xr __RUNNER_V2_EXEC__="${__ENV_ROOT__}/bin/runner-v2.sh"
declare -xr __MAKE_EXEC__="${__ENV_ROOT__}/bin/make.sh"
declare -xr __CLEANUP_EXEC__="${__ENV_ROOT__}/bin/cleanup.sh"
declare -xr __LOG__="${__ENV_ROOT__}/bin/log.sh"
declare -xr __DEBUG__="${__ENV_ROOT__}/bin/debug.sh"

. ${__ENV_ROOT__}/bin/config_profile.sh

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

  --env-file /path/to/.env:       Custom environment file.

  -d | --debug:                   Turn on debug mode, logging some useful information, nothing more.
  -s | --silent:                  Run dockers and services in background.
  -c | --cleanup:                 Clean up tmps.
  -h | --help:                    Help.

commands:
  docker:           Up and running dockers container
                    All possible containers a listed in etc/docker
                    Will create volume for corresponding containers in var/lib/<profile_name>/docker

  set:              Config profile.
                    eg: ./cli.sh -p luanphan set SOME_VAR=SOME_VALUE OTHER_VAR=OVER_VALUE

  edit:             use nvim/vim or nano to edit profile .env file
  checkconf:        printenv of current profile to screen.
  cleanconf:        clean all config of current profile.

  up:               up and running config file
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
  $__MAKE_EXEC__ profile $__PROFILE__
  __profile_config_file__="${__PROFILE_DIR__}/${__PROFILE__}/.env"
  $__DEBUG__ "EXPORT ENV from profile: ${__PROFILE__}"
  $__DEBUG__ $(cat ${__profile_config_file__})

  set -a
  . "${__profile_config_file__}"
}

while [ "$1" != "" ]; do
  case $1 in
  --env-file)
    shift
    $__DEBUG__ "EXPORT ENV from env file"
    $__DEBUG__ $(cat "${1}")
    set -a
    . "${1}"
    ;;

  -p | --profile)
    shift
    set -e
    export __PROFILE__=$1
    $__DEBUG__ -i "CLI running on profile: ${__PROFILE__}"
    apply_profile_config
    ;;

  -d | --debug)
    set -e
    export DEBUG=1
    ;;

  -s | --silent)
    set -e
    export ENV_SILENT=1
    ;;

  up)
    shift
    $__RUNNER_V2_EXEC__ $1
    exit;;

  run)
    shift
    eval $@
    exit;;

  set)
    shift
    set_env $@
    exit;;

  checkenv)
    apply_profile_config
    printenv
    exit;;

  checkconf)
    shift
    list_env $@
    exit;;

  cleanconf)
    shift
    clean_env
    exit;;

  docker)
    shift
    $__DOCKER_EXEC__ $@
    exit;;

  edit)
    shift
    edit_env
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
