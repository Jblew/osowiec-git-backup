package main

import (
	"log"

	"github.com/jblew/osowiec-git-backup/app"
)

func main() {
	err, config := GetConfig()
	if err != nil {
		log.Fatalf("Cannot get config: %v", err)
	}
	application := app.App{Config: config}
	err = application.Run()
	if err != nil {
		log.Fatalf("Cannot run app: %v", err)
	}
}
