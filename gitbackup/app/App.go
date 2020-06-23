package app

import "github.com/go-git/go-git/v5/plumbing/transport"

// App is main app runner
type App struct {
	Config       Config
	Repositories []string
	Auth         transport.AuthMethod
}
