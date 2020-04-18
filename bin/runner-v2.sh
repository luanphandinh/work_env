#!/usr/bin/env bash

. "${__ENV_ROOT__}/bin/parse_yaml.sh"
. "${__ENV_ROOT__}/bin/runner.sh"

eval "$(parse_yaml $1 global_config_)"
V2_EXECS=()

# Dockers
PROFILE=$global_config_import_profile
ENV_CONFIG=(${global_config_env[@]})
V2_EXECS+=("$(run_docker ${global_config_dockers[@]})")

# TODO: [#Config]
# Global config need to apply tmp/.env file first
# The parse again with $config_jobs to run jobs base on tmp/.env of global config
env_file="${__TMP_DIR__}/.env"
set -a
. "${env_file}"

eval "$(parse_yaml $1 config_)"
run_jobs=()
shift
if [[ -z "$@" ]]; then
  for register_job in ${config_run[@]}; do
    run_jobs+=("${register_job}")
  done
  $__LOG__ -i "Run all jobs in runner description file: ${run_jobs[@]}"
else
  for param in ${@}; do
    for register_job in ${config_run[@]}; do
      if [[ "${param}" == "${register_job}" ]]; then
        run_jobs+=("${param}")
      fi
    done
  done
  $__LOG__ -i "Run jobs: ${run_jobs[@]}"
fi

# Run jobs
for job in "${run_jobs[@]}"; do
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
