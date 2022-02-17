package app

import (
	"fmt"
	"log"

	"github.com/jblew/osowiec-git-backup/util"
)

// Run runs the app
func (app *App) Run() error {
	app.initMetrics()
	err := app.measureTimeMetric(func() error {
		return app.doPull()
	})
	app.incRunsMetric(err)
	app.pingMonitoring(err)
	app.pushMetrics()
	return err
}

func (app *App) doPull() error {
	err := app.loadRepositoryList()
	if err != nil {
		return fmt.Errorf("Cannot load repository list: %v", err)
	}

	auth, err := util.GetSSHPublicKeyFromPrivateKeyFile(app.Config.SSHPrivateKeyPath)
	if err != nil {
		return fmt.Errorf("Cannot load ssh public key from private key file: %v", err)
	}
	app.Auth = auth

	log.Printf("Repositories: %v", app.Repositories)
	err = app.pullRepositoriesSafe()
	if err != nil {
		return fmt.Errorf("Safe repository pull failed: %v", err)
	}
	return nil
}
