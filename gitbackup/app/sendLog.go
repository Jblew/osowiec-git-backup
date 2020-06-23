package app

import (
	"bytes"
	"fmt"
	"net/http"
)

// sendLog sends the log to log endpoint
func (app *App) sendLog(contents string) error {
	url := app.Config.MonitoringEndpointLog

	contentBuf := bytes.NewBuffer([]byte(contents))
	_, err := http.Post(url, "text/plain", contentBuf)
	if err != nil {
		return fmt.Errorf("Cannot send log to '%s': %v", url, err)
	}
	return nil
}
