package app

import (
	"fmt"
	"log"
	"time"
)

func (app *App) pullRepositoriesSafe() error {
	failures := 0
	fullLog := ""
	for _, repoRemoteURL := range app.Repositories {
		success, actionLog := app.pullRepo(repoRemoteURL)
		if success != true {
			failures++
		}
		log.Println(actionLog)

		fullLog += "[waiting 3 seconds]"
		time.Sleep(3 * time.Second)

		fullLog += actionLog
	}

	total := len(app.Repositories)
	succeeded := total - failures
	if failures > 0 {
		fullLog += "\n\n"
		fullLog += fmt.Sprintf("[FAILED] failures: %d, succeeeded: %d, total: %d\n", failures, succeeded, total)
	}

	err := app.sendLog(fullLog)
	if err != nil {
		return fmt.Errorf("Cannot send log: %v", err)
	}

	if failures > 0 {
		return fmt.Errorf("Pulling repositories done with %d failures out of total %d", failures, total)
	}
	return nil
}
