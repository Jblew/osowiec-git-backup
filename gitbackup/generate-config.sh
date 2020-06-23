#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"

source ./gitbackup.config.sh

if [ -z "${REPOSITORIES_LIST_ENDPOINT}" ]; then
    echo "REPOSITORIES_LIST_ENDPOINT env is not set"
    exit 1
fi

if [ -z "${REPOSITORIES_DIR}" ]; then
    echo "REPOSITORIES_DIR env is not set"
    exit 1
fi

REPOSITORIES_DIR

GO_CONFIG_FILE="${DIR}/GetAutogeneratedConfig.go"
cat >"${GO_CONFIG_FILE}" <<EOF
package main

import "gitbackup/app"

// GetAutogeneratedConfig is automatically generated file with project config
func GetAutogeneratedConfig() app.Config {
	var conf = app.Config{}
	conf.RepositoriesListEndpoint = "${REPOSITORIES_LIST_ENDPOINT}"
    conf.RepositoriesDir = "${REPOSITORIES_DIR}"

	return conf
}

EOF
