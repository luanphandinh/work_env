#!/usr/bin/env bash

usage() {
  cli_name=${0##*/}
  echo "
${0##*/}
usage: ${cli_name} [options]
options:
    -n | --name         Taking actions on profile.
                        If There is no profile, create one in etc/.

    -h | --help         Help.

commands:
    set                 Set ENV variables for profile.
    clean               Clean all ENV variables for profile.
    checkconf           List ENV variables for profile.
"
  exit 1
}

#######################################
# Apply profile name to take actions on
# Globals:
#   None
# Arguments:
#   <profile_name>
# Returns:
#   Set __PROFILE__ as <profile_name>
#######################################
profile() {
  set -e
  export __PROFILE__=$1
}

#######################################
# Apply ENV value to __PROFILE__
# Globals:
#   __PROFILE__
# Arguments:
#   pairs of values
#   ex: FOO=BAR
# Returns:
#   append FOO=BAR to $__PROFILE_DIR__/$__PROFILE__/.env file
#######################################
set_env() {
  env="${__PROFILE_DIR__}/${__PROFILE__}/.env"
  if [[ ! -d "${__PROFILE_DIR__}/${__PROFILE__}" ]]; then
    mkdir "${__PROFILE_DIR__}/${__PROFILE__}"
  fi

  VARIBABLES=$@
  for VARIBABLE in ${VARIBABLES}; do
    echo "${VARIBABLE}" >> "${env}"
  done
}

clean_env() {
  env="${__PROFILE_DIR__}/${__PROFILE__}/.env"
  if [[ ! -d "${__PROFILE_DIR__}/${__PROFILE__}" ]]; then
    return
  fi

  > "${env}"
}

list_env() {
  if [[ "--all" = $1 || "-a" = $1 ]]; then
    echo "=============== DEFAULT_CLI_CONFIG ==============="
    cat "${__DOCKER_DIR__}/.PORT"
  fi

  echo "=============== ${__PROFILE__}_CONFIG ==============="
  cat "${__PROFILE_DIR__}/${__PROFILE__}/.env"
}

while [ "$1" != "" ]; do
  case $1 in
  -n | --name)
    shift
    profile $1
    ;;

  set)
    shift
    set_env $@
    exit;;

  clean)
    shift
    clean_env $@
    exit;;

  checkconf)
    shift
    list_env $@
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
