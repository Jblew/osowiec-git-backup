package app

import (
	"log"

	"github.com/jblew/osowiec-git-backup/util"
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
	pullHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "gitbackup_pull_time",
		Help:    "Pull time histogram of all repositories",
		Buckets: prometheus.LinearBuckets(20, 5, 5),
	})
	pullCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gitbackup_pulls",
		Help: "Pull counter",
	}, []string{"status", "type"})
	branchesCountGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gitbackup_branches_total",
		Help: "Total number of branches",
	})
	commitCountGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gitbackup_commits_total",
		Help: "Total number of commits",
	})
	retriesCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "gitbackup_retries_total",
		Help: "Total number of retries",
	})
)

func (app *App) initMetrics() {
	prometheus.MustRegister(runsCounter, timeGauge, repositoryCountGauge, pullHistogram, pullCounter, branchesCountGauge, commitCountGauge, retriesCounter)
}

func (app *App) pushMetrics() {
	gathered, _ := prometheus.DefaultGatherer.Gather()
	log.Printf("Pushing the following metrics: %+v", gathered)

	if app.Config.PrometheusPushGatewayURL == "" {
		log.Printf("PrometheusPushGatewayURL not configured via env (consult Dockerfile). Skipping prometheus metrics pushing")
		return
	}
	err := prometheusPush.New(app.Config.PrometheusPushGatewayURL, app.Config.PrometheusJobName).Gatherer(prometheus.DefaultGatherer).Push()
	if err != nil {
		log.Printf("Error while pushing prometheus metrics: %v", err)
	} else {
		log.Printf("Successfuly pushed metrics")
	}
}

func (app *App) measureRunTimeMetric(f func() error) error {
	duration, err := util.MeasureDuration(f)
	if err != nil {
		timeGauge.With(prometheus.Labels{"status": "error"}).Set(duration.Seconds())
	} else {
		timeGauge.With(prometheus.Labels{"status": "success"}).Set(duration.Seconds())
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

func (app *App) incBranchesMetric(count int) {
	branchesCountGauge.Add(float64(count))
}

func (app *App) incCommitsMetric(count int) {
	commitCountGauge.Add(float64(count))
}

func (app *App) incRetriesMetric() {
	retriesCounter.Inc()
}

func (app *App) incPullsMetricSuccess(typeName string) {
	pullCounter.With(prometheus.Labels{"status": "success", "type": typeName}).Inc()
}

func (app *App) incPullsMetricFailure() {
	pullCounter.With(prometheus.Labels{"status": "failure", "type": "error"}).Inc()
}

func (app *App) measurePullTimeMetric(f func() error) error {
	duration, err := util.MeasureDuration(f)
	pullHistogram.Observe(duration.Seconds())
	return err
}

func (app *App) setRepositoriesCountMetric(len int) {
	repositoryCountGauge.Set(float64(len))
}
