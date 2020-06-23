package app

import "log"

// Run runs the app
func Run(config Config) error {
	app := App{Config: config}
	err := app.loadRepositoryList()
	if err != nil {
		return err
	}

	log.Printf("Repositories: %v", app.Repositories)

	return nil
}
