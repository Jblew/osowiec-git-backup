#!/usr/bin/env bash
PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export DIST_DIR="${PROJECT_DIR}/dist"
export REPOSITORIES_LIST_ENDPOINT="https://raw.githubusercontent.com/Jblew/osowiec-git-backup/master/repositories.lst"
export REPOSITORIES_DIR=""
export SSH_PRIVATE_KEY_PATH=""
export MONITORING_ENDPOINT_LOG=""
export MONITORING_ENDPOINT_PING_SUCCESS=""
export MONITORING_ENDPOINT_PING_FAILURE=""
# I use https://healthchecks.io/ for ping monitoring
