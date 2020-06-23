package app

import "log"

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
	}

	// TODO publish fullLog to log endpoint
	// if failures > 0 mark as error

	return nil
}
