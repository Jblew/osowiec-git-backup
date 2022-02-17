package app

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	prometheusPush "github.com/prometheus/client_golang/prometheus/push"
)

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

func (app *App) initMetrics() {
	prometheus.MustRegister(runsCounter, timeGauge, repositoryCountGauge)
}

func (app *App) pushMetrics() {
	gathered, _ := prometheus.DefaultGatherer.Gather()
	fmt.Printf("Pushing the following metrics: %+v", gathered)

	if app.Config.PrometheusPushGatewayURL == "" {
		fmt.Printf("PrometheusPushGatewayURL not configured via env (consult Dockerfile). Skipping prometheus metrics pushing")
		return
	}
	err := prometheusPush.New(app.Config.PrometheusPushGatewayURL, app.Config.PrometheusJobName).Gatherer(prometheus.DefaultGatherer).Push()
	if err != nil {
		fmt.Printf("Error while pushing prometheus metrics: %v", err)
	} else {
		fmt.Printf("Successfuly pushed metrics")
	}
}

func (app *App) measureTimeMetric(f func() error) error {
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

func (app *App) incRunsMetric(err error) {
	if err != nil {
		runsCounter.With(prometheus.Labels{"status": "failure"}).Inc()
	} else {
		runsCounter.With(prometheus.Labels{"status": "failure"}).Inc()
	}
}

func (app *App) setRepositoriesCountMetric(len int) {
	repositoryCountGauge.Set(float64(len))
}
