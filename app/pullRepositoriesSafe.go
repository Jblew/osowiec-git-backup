package app

import (
	"fmt"
	"time"
)

func (app *App) pullRepositoriesSafe() error {
	failures := []error{}
	for _, repoRemoteURL := range app.Repositories {
		err := app.pullRepo(repoRemoteURL)
		if err != nil {
			failures = append(failures, err)
		}
		time.Sleep(2 * time.Second)
	}

	if len(failures) > 0 {
		return fmt.Errorf("Pulling repositories done with %d failures out of total %d", len(failures), len(app.Repositories))
	}
	return nil
}
