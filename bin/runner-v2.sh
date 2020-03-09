#!/usr/bin/env bash

. "${__ENV_ROOT__}/bin/parse_yaml.sh"
. "${__ENV_ROOT__}/bin/runner.sh"

eval "$(parse_yaml $1 config_)"
V2_EXECS=()

# Dockers
PROFILE=$config_profile
ENV_CONFIG=(${config_env[@]})
V2_EXECS+=("$(run_docker ${config_dockers[@]})")

# All the services
for service in "${config_services[@]}"; do
  service_profile="config_service_${service}_profile"
  service_env="config_service_${service}_env[*]"
  service_path="config_service_${service}_path"
  service_run="config_service_${service}_run"

  PROFILE=${!service_profile}
  ENV_CONFIG=(${!service_env})
  SERVICE_PATH=${!service_path}
  RUN=${!service_run}

  cli=$(cli_command "${service}")
  cmd="(cd ${SERVICE_PATH} && ${RUN})"
  if [[ $ENV_SILENT == 1 ]]; then
    cmd+=" &"
  fi
  V2_EXECS+=("${cli} run ${cmd}")
done

for service in "${V2_EXECS[@]}";
do
  echo "${service}"
  $service &
done

wait
