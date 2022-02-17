package app

import (
	"fmt"
	"gitbackup/util"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	prometheusPush "github.com/prometheus/client_golang/prometheus/push"
)

// Run runs the app
func Run(config Config) error {
	initMetrics()
	app := App{Config: config}
	runFunc := func() error {
		return app.doPull()
	}
	err := measureTimeMetric(runFunc)

	if err != nil {
		runsCounter.With(prometheus.Labels{"status": "failure"}).Inc()
		pingErr := app.pingMonitoringFailure()
		if pingErr != nil {
			fmt.Printf("Pull failed with [%v] and cannot ping monitoring failure [%v]", err, pingErr)
		}
	} else {
		runsCounter.With(prometheus.Labels{"status": "success"}).Inc()
		pingErr := app.pingMonitoringSuccess()
		if pingErr != nil {
			fmt.Printf("Pull succeeded but cannot ping monitoring success [%v]", pingErr)
		}
	}
	pushMetrics(config)

	return nil
}

func (app *App) doPull() error {
	err := app.loadRepositoryList()
	if err != nil {
		return fmt.Errorf("Cannot load repository list: %v", err)
	}
	repositoryCountGauge.Set(float64(len(app.Repositories)))

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

var (
	runsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gitbackup_runs_count",
		Help: "Count of runs of gitbackup",
	}, []string{"status"})

	timeGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gitbackup_run_time",
		Help: "Time of running gitbackup",
	}, []string{"status"})

	repositoryCountGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gitbackup_repository_count",
		Help: "Number of all repositories",
	})
)

func initMetrics() {
	prometheus.MustRegister(runsCounter, timeGauge, repositoryCountGauge)
}

func pushMetrics(config Config) {
	if config.PrometheusPushGatewayURL == "" {
		fmt.Printf("PrometheusPushGatewayURL not configured via env (consult Dockerfile). Skipping prometheus metrics pushing")
		return
	}
	err := prometheusPush.New(config.PrometheusPushGatewayURL, config.PrometheusJobName).Gatherer(prometheus.DefaultGatherer).Push()
	if err != nil {
		fmt.Printf("Error while pushing prometheus metrics: %v", err)
	}
}

func measureTimeMetric(f func() error) error {
	start := time.Now()
	err := f()
	ellapsedSeconds := time.Since(start)
	if err != nil {
		timeGauge.With(prometheus.Labels{"status": "error"}).Set(ellapsedSeconds.Seconds())
	} else {
		timeGauge.With(prometheus.Labels{"status": "success"}).Set(ellapsedSeconds.Seconds())
	}
	return err
}
