#!/usr/bin/env bash

DOCKERS_PATTERN="\[dockers\:.*\]"
PROFILE_PATTERN="\[profile\:.*\]"
ENV_PATTERN="\[env\]"

DOCKERS=""
PROFILE=""
ENV_CONFIG=""

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
    if [[ "$1" != ${LOCK_ON} ]]; then
      echo $1
    fi
  fi
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
      lock "${ENV_PATTERN}"
    fi

    pushConfig $line
  done < "$configFile"
}

run $@
