package app

import (
	"fmt"
	"gitbackup/util"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

// Run runs the app
func Run(config Config) error {
	app := App{Config: config}
	err := app.doPull()
	if err != nil {
		pingErr := app.pingMonitoringFailure()
		if pingErr != nil {
			return fmt.Errorf("Pull failed with [%v] and cannot ping monitoring failure [%v]", err, pingErr)
		}
		return err
	}

	pingErr := app.pingMonitoringSuccess()
	if pingErr != nil {
		return fmt.Errorf("Pull succeeded but cannot ping monitoring success [%v]", pingErr)
	}

	return nil
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

	auth.HostKeyCallback = MakeEmptyHostkeyCallback()
	app.Auth = auth

	log.Printf("Repositories: %v", app.Repositories)
	err = app.pullRepositoriesSafe()
	if err != nil {
		return fmt.Errorf("Safe repository pull failed: %v", err)
	}
	return nil
}

func MakeEmptyHostkeyCallback() ssh.HostKeyCallback {
	// allows all known hosts
	return func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

}
