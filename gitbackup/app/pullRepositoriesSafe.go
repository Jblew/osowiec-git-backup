package app

import (
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

	// TODO publish fullLog to log endpoint
	// if failures > 0 mark as error

	return nil
}
