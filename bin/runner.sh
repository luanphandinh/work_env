#!/usr/bin/env bash

DOCKERS_PATTERN="\[dockers\:.*\]"
PROFILE_PATTERN="\[profile\:.*\]"
ENV_PATTERN="\[env\]"

DOCKERS=""
PROFILE=""
ENV_CONFIG=()

LOCK_ON=""

unlock() {
  LOCK_ON=""
}

lock() {
  LOCK_ON=$1
}

isLock() {
  [[ ! -z "${LOCK_ON}" ]]
  return
}

pushConfig() {
  if isLock; then
    if [[ ! $1 == ${LOCK_ON} ]]; then
      if [[ "${LOCK_ON}" == "${ENV_PATTERN}" ]]; then
        ENV_CONFIG+=("$1")
      fi
    fi
  fi
}

compile() {
  cmd="${__ENV_ROOT__}/cli.sh"

  if [[ ! -z "${PROFILE}" ]]; then
    cmd+=" -p ${PROFILE}"
  fi

  if [[ ! -z "${__TMP_DIR__}" ]]; then
    cmd+=" --env-file ${__TMP_DIR__}/.env"
    > "${__TMP_DIR__}/.env"
    for conf in "${ENV_CONFIG[@]}";
    do
      echo "${conf}" >> "${__TMP_DIR__}/.env"
    done
  fi

  if [[ ! -z "${DOCKERS}" ]]; then
    cmd+=" docker run ${DOCKERS}"
  fi

  echo "${cmd}" | tr -s " "
}

run() {
  configFile=$1
  while read -r line;
  do
    if [[ "${line}" =~ ${DOCKERS_PATTERN} ]]; then
      DOCKERS="${line/\[dockers\:/}"
      DOCKERS="${DOCKERS/]/}"
      unlock
    fi

    if [[ "${line}" =~ ${PROFILE_PATTERN} ]]; then
      PROFILE="${line/\[profile\:/}"
      PROFILE="${PROFILE/]/}"
      unlock
    fi

    if [[ "${line}" =~ ${ENV_PATTERN} ]]; then
      lock ${ENV_PATTERN}
    fi

    pushConfig "${line}"
  done < "$configFile"

  echo "$(compile)"
  $(compile)
}

run $@
