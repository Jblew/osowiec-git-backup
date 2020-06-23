package app

import (
	"fmt"
	"net/http"
)

// pingMonitoringSuccess send success ping to the configured monitoring endpoint
func (app *App) pingMonitoringSuccess() error {
	return pingMonitoring(app.Config.MonitoringEndpointPingSuccess)
}

// pingMonitoringFailure send failure ping to the configured monitoring endpoint
func (app *App) pingMonitoringFailure() error {
	return pingMonitoring(app.Config.MonitoringEndpointPingFailure)
}

func pingMonitoring(url string) error {
	_, err := http.Head(url)
	if err != nil {
		return fmt.Errorf("Cannot send ping to '%s': %v", url, err)
	}
	return nil
}
