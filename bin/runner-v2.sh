#!/usr/bin/env bash

. "${__ENV_ROOT__}/bin/parse_yaml.sh"
. "${__ENV_ROOT__}/bin/runner.sh"

eval "$(parse_yaml $1 config_)"
V2_EXECS=()

# Dockers
PROFILE=$config_import_profile
ENV_CONFIG=(${config_env[@]})
V2_EXECS+=("$(run_docker ${config_dockers[@]})")

# All the services
for job in "${config_run[@]}"; do
  job_profile="config_jobs_${job}_import_profile"
  job_env="config_jobs_${job}_env[*]"
  job_path="config_jobs_${job}_path"
  job_run="config_jobs_${job}_run"

  # VARIABLES that runner.sh uses.
  # FUNCS: cli_command, run_docker.
  PROFILE=${!job_profile}
  ENV_CONFIG=(${!job_env})
  SERVICE_PATH=${!job_path}
  RUN=${!job_run}

  cli=$(cli_command "${job}")
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
