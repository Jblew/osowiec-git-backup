package app

import (
	"fmt"
	"gitbackup/util"
	"log"
)

// Run runs the app
func Run(config Config) error {
	app := App{Config: config}
	err := app.loadRepositoryList()
	if err != nil {
		return fmt.Errorf("Cannot load repository list: %v", err)
	}

	auth, err := util.GetSSHPublicKeyFromPrivateKeyFile(config.SSHPrivateKeyPath)
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
