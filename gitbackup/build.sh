#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

source ./gitbackup.config.sh

./generate-config.sh
./test.sh

go build -o "${DIST_DIR}/gitbackup" *.go
