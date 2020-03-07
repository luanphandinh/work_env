#!/usr/bin/env bash

. "${__ENV_ROOT__}/bin/parse_yaml.sh"
. "${__ENV_ROOT__}/bin/runner.sh"

eval "$(parse_yaml $1 config_)"

$(run_docker "${config_dockers[*]}")
