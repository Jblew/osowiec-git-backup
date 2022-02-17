package app

import (
	"fmt"
	"log"
	"net/http"
)

func (app *App) pingMonitoring(err error) {
	if err != nil {
		app.pingMonitoringFailure(err)
	} else {
		app.pingMonitoringSuccess()
	}
}

// pingMonitoringSuccess send success ping to the configured monitoring endpoint
func (app *App) pingMonitoringSuccess() {
	err := doPingMonitoring(app.Config.MonitoringEndpointPingSuccess)
	if err != nil {
		fmt.Printf("Pull succeeded but cannot ping monitoring success [%v]", err)
	}
}

// pingMonitoringFailure send failure ping to the configured monitoring endpoint
func (app *App) pingMonitoringFailure(appErr error) {
	pingErr := doPingMonitoring(app.Config.MonitoringEndpointPingFailure)
	if pingErr != nil {
		fmt.Printf("Pull failed with [%v] and cannot ping monitoring failure [%v]", appErr, pingErr)
	}
}

func doPingMonitoring(url string) error {
	if url == "" {
		log.Printf("Monitoring endpoints for ping not specified. Skipping ping sending")
		return nil
	}
	_, err := http.Head(url)
	if err != nil {
		return fmt.Errorf("Cannot send ping to '%s': %v", url, err)
	}
	log.Printf("Monitoring ping sent to '%s'\n", url)
	return nil
}
