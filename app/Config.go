package app

// Config is App configuration
type Config struct {
	RepositoriesListFile          string
	RepositoriesDir               string
	SSHPrivateKeyPath             string
	MonitoringEndpointPingSuccess string
	MonitoringEndpointPingFailure string
	PrometheusPushGatewayURL      string
	PrometheusJobName             string
}
