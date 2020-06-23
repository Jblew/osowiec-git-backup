#!/usr/bin/env bash
PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export DIST_DIR="${PROJECT_DIR}/dist"
export REPOSITORIES_LIST_ENDPOINT="https://raw.githubusercontent.com/Jblew/osowiec-git-backup/master/repositories.lst"
export REPOSITORIES_DIR="/mnt/repos"
