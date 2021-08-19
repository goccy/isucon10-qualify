package main

import (
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	newrelicApp *newrelic.Application
)

func init() {
	os.Setenv("NEW_RELIC_LICENSE_KEY", "")
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("ISUCON10"),
		newrelic.ConfigFromEnvironment(),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		panic(err)
	}
	newrelicApp = app
}
