#!/usr/bin/env bash

usage() {
  cli_name=${0##*/}
  echo "
${0##*/}
usage: ${cli_name} [options]
options:
  profile      Create new profile in /etc/profile
  *            Help
"
  exit 1
}

make_profile() {
  profile=$1
  __dir__="${__ENV_ROOT__}/etc/profile/${profile}"
  if [[ ! -d "${__dir__}" ]]; then
    $__DEBUG__ "Create new profile: ${__PROFILE__}"
    mkdir "${__dir__}"
  fi

  __profile_config_file__="${__dir__}/.env"
  if [[ ! -d "${__profile_config_file__}" ]]; then
    touch "${__profile_config_file__}"
  fi
}

while [ "$1" != "" ]; do
  case $1 in
  profile)
    shift
    make_profile $1
    exit
    ;;

  *)
    usage
    exit 1
    ;;
  esac
  shift
done
