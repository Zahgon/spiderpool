// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package metric

import (
	"context"
	"net/http"
	"time"

	api "go.opentelemetry.io/otel/metric"
)

const debugMetrics = "debug-metrics"

var (
	// meter is a global creator of metric instruments.
	meter, debugLevelMeter api.Meter
	// globalEnableMetric determines whether to use metric or not
	globalEnableMetric bool
)

// InitMetric will set up meter with the input param(required) and create a prometheus exporter.
func InitMetric(ctx context.Context, meterName string, enableMetric, enableDebugLevelMetrics bool) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

// newMetricInt64Counter will create otel Int64Counter metric.
// The first param metricName is required and the second param is optional.
func newMetricInt64Counter(metricName string, description string, isDebugLevel bool) (api.Int64Counter, error) {
	_ = "STUB: not implemented"
	return *new(api.Int64Counter), nil
}

// newMetricFloat64Histogram will create otel Float64Histogram metric.
// The first param metricName is required and the second param is optional.
// Notice: if you want to match the quantile {0.1, 0.3, 0.5, 1, 3, 5, 7, 10, 15}, please let the metric name match regex "*_histogram",
// otherwise it will match the  otel default quantile.
func newMetricFloat64Histogram(metricName string, description string, isDebugLevel bool) (api.Float64Histogram, error) {
	_ = "STUB: not implemented"
	return *new(api.Float64Histogram), nil
}

// newMetricFloat64Gauge will create otel Float64Gauge metric.
// The first param metricName is required and the second param is optional.
func newMetricFloat64Gauge(metricName string, description string, isDebugLevel bool) (api.Float64ObservableGauge, error) {
	_ = "STUB: not implemented"
	return *new(api.Float64ObservableGauge), nil
}

// newMetricInt64Gauge will create otel Int64Gauge metric.
// The first param metricName is required and the second param is optional.
func newMetricInt64Gauge(metricName string, description string, isDebugLevel bool) (api.Int64ObservableGauge, error) {
	_ = "STUB: not implemented"
	return *new(api.Int64ObservableGauge), nil
}

var _ TimeRecorder = &timeRecorder{}

// timeRecorder owns a field to record start time.
type timeRecorder struct {
	startTime time.Time
}

// TimeRecorder will help you to compute time duration.
type TimeRecorder interface {
	SinceInSeconds() float64
}

// NewTimeRecorder will create TimeRecorder and record the current time.
func NewTimeRecorder() TimeRecorder { _ = "STUB: not implemented"; return *new(TimeRecorder) }

// SinceInSeconds returns the duration of time since the start time as a float64.
func (t *timeRecorder) SinceInSeconds() float64 { _ = "STUB: not implemented"; return 0 }
