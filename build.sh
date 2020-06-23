#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

source ./project.config.sh

echo "# Cleaning"
rm -rf "${DIST_DIR}"
echo "# Cleaning done"
echo ""

echo "# Building gitbackup"
./gitbackup/build.sh
echo "# Gitbackup build done"
echo ""

