package app

// Config is App configuration
type Config struct {
	RepositoriesListEndpoint      string
	RepositoriesDir               string
	SSHPrivateKeyPath             string
	MonitoringEndpointLog         string
	MonitoringEndpointPingSuccess string
	MonitoringEndpointPingFailure string
}
