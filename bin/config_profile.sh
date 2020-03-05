#!/usr/bin/env bash

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

edit_env() {
  if [[ $(nvim --version) ]]; then
    nvim "${__PROFILE_DIR__}/${__PROFILE__}/.env"
    return
  fi

  if [[ $(vim --version) ]]; then
    vim "${__PROFILE_DIR__}/${__PROFILE__}/.env"
    return
  fi

  if [[ $(nano --version) ]]; then
    nano "${__PROFILE_DIR__}/${__PROFILE__}/.env"
    return
  fi
}
