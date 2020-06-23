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

		fullLog += actionLog
		fullLog += "[waiting 3 seconds]\n\n\n"
		time.Sleep(3 * time.Second)
	}

	total := len(app.Repositories)
	succeeded := total - failures
	if failures > 0 {
		fullLog += "\n\n"
		fullLog += fmt.Sprintf("[FAILED] failures: %d, succeeeded: %d, total: %d\n", failures, succeeded, total)
	}

	fullLog += fmt.Sprintf("\n\n[SUCCESS] %d synchronized, 0 failures\n", total)

	err := app.sendLog(fullLog)
	if err != nil {
		return fmt.Errorf("Cannot send log: %v", err)
	}

	if failures > 0 {
		return fmt.Errorf("Pulling repositories done with %d failures out of total %d", failures, total)
	}
	return nil
}
