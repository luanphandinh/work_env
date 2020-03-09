#!/usr/bin/env bash

DOCKERS_PATTERN="[dockers:"
PROFILE_PATTERN="[profile:"
ENV_PATTERN="[env]"
EXEC_PATTERN="[exec:"

DOCKERS=""
PROFILE=""
ENV_CONFIG=()
EXEC_CMDS=()
SERVICES=()

LOCK_ON=""

unlock() {
  LOCK_ON=""
}

lock() {
  LOCK_ON="${1}"
}

isLock() {
  [[ ! -z "${LOCK_ON}" ]]
  return
}

pushConfig() {
  if isLock; then
    if [[ ! "${1}" =~ "${LOCK_ON}" ]]; then
      if [[ "${LOCK_ON}" == "${ENV_PATTERN}" ]]; then
        ENV_CONFIG+=("$1")
      fi

      if [[ "${LOCK_ON}" =~ "${EXEC_PATTERN}" ]]; then
        EXEC_CMDS+=("$1")
      fi
    fi
  fi
}

cli_command() {
  local name=$1
  local cmd="${__ENV_ROOT__}/cli.sh"

  if [[ ! -z "${PROFILE}" ]]; then
    cmd+=" -p ${PROFILE}"
  fi

  env_file="${__TMP_DIR__}/.env"
  if [[ ! -z "${name}" ]]; then
    env_file+=".${name}"
  fi

  cmd+=" --env-file ${env_file}"
  > "${env_file}"
  for conf in "${ENV_CONFIG[@]}";
  do
    echo "${conf}" >> "${env_file}"
  done

  echo "${cmd}" | tr -s " "
}

exec_services() {
  if [[ ${#EXEC_CMDS[@]} -gt 0 ]]; then
    local cli=$(cli_command)

    local ITER=0
    for conf in "${EXEC_CMDS[@]}";
    do
      cmd+=" ${conf}"
      ITER=$(expr $ITER + 1)
      if [[ $ITER -lt ${#EXEC_CMDS[@]} ]]; then
        cmd+=" &&"
      fi
    done

    if [[ $DEBUG != 1 ]]; then
      cmd+=" &"
    fi

    echo "${cli} run ${cmd}" | tr -s " "
  fi
}

run_docker() {
  local names=$@
  local cmd=$(cli_command)
  if [[ ! -z "${names}" ]]; then
    cmd+=" docker run ${names}"
    echo "${cmd}" | tr -s " "
  fi
}

run() {
  configFile=$1
  while read -r line;
  do
    if [[ ! $line = *[!\ ]* ]]; then
      continue
    fi

    if [[ $line =~ ^\#.*|^\/\/.* ]];then
      continue
    fi

    if [[ "${line}" =~ "${DOCKERS_PATTERN}" ]]; then
      DOCKERS="${line/\[dockers\:/}"
      DOCKERS="${DOCKERS/]/}"
      unlock
    fi

    if [[ "${line}" =~ "${PROFILE_PATTERN}" ]]; then
      PROFILE="${line/\[profile\:/}"
      PROFILE="${PROFILE/]/}"
      unlock
    fi

    if [[ "${line}" =~ "${ENV_PATTERN}" ]]; then
      lock ${ENV_PATTERN}
    fi

    if [[ "${line}" =~ "${EXEC_PATTERN}" ]]; then
      if [[ isLock ]]; then
        SERVICES+=("$(exec_services)")
        EXEC_CMDS=()
      fi
      lock ${line}
    fi

    pushConfig "${line}"
  done < "$configFile"

  SERVICES+=("$(exec_services)")

  $__LOG__ -i "Run and send to background: $(run_docker)"
  $(run_docker "${DOCKERS}") &

  for service in "${SERVICES[@]}";
  do
    $service &
  done

  wait
}

# run $@
