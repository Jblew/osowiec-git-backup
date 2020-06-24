package app

import (
	"fmt"
	"log"
	"time"
)

func (app *App) pullRepositoriesSafe() error {
	failures := []error{}
	fullLog := ""
	for _, repoRemoteURL := range app.Repositories {
		actionLog, err := app.pullRepo(repoRemoteURL)
		if err != nil {
			failures = append(failures, err)
		}
		log.Println(actionLog)

		fullLog += actionLog
		fullLog += "[waiting 2 seconds]\n\n"
		time.Sleep(2 * time.Second)
	}

	fullLog += "\n\n"

	footer := formatFooter(app.Repositories, failures)
	fullLog += footer
	log.Println(footer)

	err := app.sendLog(fullLog)
	if err != nil {
		return fmt.Errorf("Cannot send log: %v", err)
	}

	if len(failures) > 0 {
		return fmt.Errorf("Pulling repositories done with %d failures out of total %d", len(failures), len(app.Repositories))
	}
	return nil
}

func formatFooter(repositories []string, failures []error) string {
	numTotal := len(repositories)
	numFailures := len(failures)
	numSucceeded := numTotal - numFailures

	if len(failures) > 0 {
		out := fmt.Sprintf("[FAILED] failures: %d, succeeeded: %d, total: %d\n", numFailures, numSucceeded, numTotal)
		for _, failureErr := range failures {
			out += fmt.Sprintf("  %v\n", failureErr)
		}
		return out
	}
	return fmt.Sprintf("\n\n[SUCCESS] %d synchronized, 0 failures\n", numTotal)

}
