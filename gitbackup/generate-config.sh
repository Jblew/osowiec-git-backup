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

if [ -z "${SSH_PRIVATE_KEY_PATH}" ]; then
    echo "SSH_PRIVATE_KEY_PATH env is not set"
    exit 1
fi

if [ -z "${MONITORING_ENDPOINT_LOG}" ]; then
    echo "MONITORING_ENDPOINT_LOG env is not set"
    exit 1
fi

if [ -z "${MONITORING_ENDPOINT_PING_SUCCESS}" ]; then
    echo "MONITORING_ENDPOINT_PING_SUCCESS env is not set"
    exit 1
fi

if [ -z "${MONITORING_ENDPOINT_PING_FAILURE}" ]; then
    echo "MONITORING_ENDPOINT_PING_FAILURE env is not set"
    exit 1
fi

GO_CONFIG_FILE="${DIR}/GetAutogeneratedConfig.go"
cat >"${GO_CONFIG_FILE}" <<EOF
package main

import "gitbackup/app"

// GetAutogeneratedConfig is automatically generated file with project config
func GetAutogeneratedConfig() app.Config {
	var conf = app.Config{}
	conf.RepositoriesListEndpoint = "${REPOSITORIES_LIST_ENDPOINT}"
    conf.RepositoriesDir = "${REPOSITORIES_DIR}"
    conf.SSHPrivateKeyPath = "${SSH_PRIVATE_KEY_PATH}"
    conf.MonitoringEndpointLog="${MONITORING_ENDPOINT_LOG}"
    conf.MonitoringEndpointPingSuccess="${MONITORING_ENDPOINT_PING_SUCCESS}"
    conf.MonitoringEndpointPingFailure="${MONITORING_ENDPOINT_PING_FAILURE}"

	return conf
}

EOF
