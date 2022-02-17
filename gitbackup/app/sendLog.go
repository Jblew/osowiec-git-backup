package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// sendLog sends the log to log endpoint
func (app *App) sendLog(contents string) error {
	url := app.Config.MonitoringEndpointLog
	if url == "" {
		log.Printf("MonitoringEndpointLog URL not specified, sending logs skipped\n\n")
		return nil
	}

	contentBuf := bytes.NewBuffer([]byte(contents))
	response, err := http.Post(url, "text/plain", contentBuf)
	if err != nil {
		return fmt.Errorf("Cannot send log to '%s': %v", url, err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	responseStr := string(responseData)

	log.Printf("Sent log. Response: %s\n\n", responseStr)

	if response.StatusCode != 200 {
		return fmt.Errorf("sendLog resonse code is not 200: %d, response: %s", response.StatusCode, responseStr)
	}

	return nil
}
