// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package metric

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	api "go.opentelemetry.io/otel/metric"

	"github.com/spidernet-io/spiderpool/pkg/lock"
	"github.com/spidernet-io/spiderpool/pkg/podownercache"
)

const (
	metricPrefix = "spiderpool_"
	debugPrefix  = "debug_"
)

const (
	// spiderpool agent ipam allocation metrics name
	ipamAllocationCountsName                     = metricPrefix + "ipamAllocationCountsName"
	ipamAllocationFailureCountsName              = metricPrefix + "ipamAllocationFailureCountsName"
	ipamAllocationUpdateIPPoolConflictCountsName = metricPrefix + "ipamAllocationUpdateIPPoolConflictCountsName"
	ipamAllocationErrInternalCountsName          = metricPrefix + "ipamAllocationErrInternalCountsName"
	ipamAllocationErrNoAvailablePoolCountsName   = metricPrefix + "ipamAllocationErrNoAvailablePoolCountsName"
	ipamAllocationErrRetriesExhaustedCountsName  = metricPrefix + "ipamAllocationErrRetriesExhaustedCountsName"
	ipamAllocationErrIPUsedOutCountsName         = metricPrefix + "ipamAllocationErrIPUsedOutCountsName"

	ipamAllocationAverageDurationSecondsName = metricPrefix + "ipamAllocationAverageDurationSecondsName"
	ipamAllocationMaxDurationSecondsName     = metricPrefix + "ipamAllocationMaxDurationSecondsName"
	ipamAllocationMinDurationSecondsName     = metricPrefix + "ipamAllocationMinDurationSecondsName"
	ipamAllocationLatestDurationSecondsName  = metricPrefix + "ipamAllocationLatestDurationSecondsName"
	ipamAllocationDurationSecondsName        = metricPrefix + "ipamAllocationDurationSecondsName"

	ipamAllocationAverageLimitDurationSecondsName = metricPrefix + "ipamAllocationAverageLimitDurationSecondsName"
	ipamAllocationMaxLimitDurationSecondsName     = metricPrefix + "ipamAllocationMaxLimitDurationSecondsName"
	ipamAllocationMinLimitDurationSecondsName     = metricPrefix + "ipamAllocationMinLimitDurationSecondsName"
	ipamAllocationLatestLimitDurationSecondsName  = metricPrefix + "ipamAllocationLatestLimitDurationSecondsName"
	ipamAllocationLimitDurationSecondsName        = metricPrefix + "ipamAllocationLimitDurationSecondsName"

	// spiderpool agent ipam release metrics name
	ipamReleaseCountsName                     = metricPrefix + "ipamReleaseCountsName"
	ipamReleaseFailureCountsName              = metricPrefix + "ipamReleaseFailureCountsName"
	ipamReleaseUpdateIPPoolConflictCountsName = metricPrefix + "ipamReleaseUpdateIPPoolConflictCountsName"
	ipamReleaseErrInternalCountsName          = metricPrefix + "ipamReleaseErrInternalCountsName"
	ipamReleaseErrRetriesExhaustedCountsName  = metricPrefix + "ipamReleaseErrRetriesExhaustedCountsName"

	ipamReleaseAverageDurationSecondsName = metricPrefix + "ipamReleaseAverageDurationSecondsName"
	ipamReleaseMaxDurationSecondsName     = metricPrefix + "ipamReleaseMaxDurationSecondsName"
	ipamReleaseMinDurationSecondsName     = metricPrefix + "ipamReleaseMinDurationSecondsName"
	ipamReleaseLatestDurationSecondsName  = metricPrefix + "ipamReleaseLatestDurationSecondsName"
	ipamReleaseDurationSecondsName        = metricPrefix + "ipamReleaseDurationSecondsName"

	ipamReleaseAverageLimitDurationSecondsName = metricPrefix + "ipamReleaseAverageLimitDurationSecondsName"
	ipamReleaseMaxLimitDurationSecondsName     = metricPrefix + "ipamReleaseMaxLimitDurationSecondsName"
	ipamReleaseMinLimitDurationSecondsName     = metricPrefix + "ipamReleaseMinLimitDurationSecondsName"
	ipamReleaseLatestLimitDurationSecondsName  = metricPrefix + "ipamReleaseLatestLimitDurationSecondsName"
	ipamReleaseLimitDurationSecondsName        = metricPrefix + "ipamReleaseLimitDurationSecondsName"

	// spiderpool controller IP GC metrics name
	ipGCCCountsName       = metricPrefix + "ipGCCCountsName"
	ipGCFailureCountsName = metricPrefix + "ipGCFailureCountsName"

	// spiderpool IPPool and Subnet metrics and these include some debug level metrics
	totalIPPoolCountsName                = metricPrefix + "totalIPPoolCountsName"
	ippoolTotalIPCountsName              = metricPrefix + debugPrefix + "ippoolTotalIPCountsName"
	ippoolAvailableIPCountsName          = metricPrefix + debugPrefix + "ippoolAvailableIPCountsName"
	totalSubnetCountsName                = metricPrefix + "totalSubnetCountsName"
	subnetIPPoolCountsName               = metricPrefix + debugPrefix + "subnetIPPoolCountsName"
	subnetTotalIPCountsName              = metricPrefix + debugPrefix + "subnetTotalIPCountsName"
	subnetAvailableIPCountsName          = metricPrefix + debugPrefix + "subnetAvailableIPCountsName"
	autoPoolWaitedForAvailableCountsName = metricPrefix + debugPrefix + "autoPoolWaitedForAvailableCountsName"
)

var (
	// ipam allocation metrics in spiderpool-agent
	IpamAllocationTotalCounts                   api.Int64Counter
	IpamAllocationFailureCounts                 api.Int64Counter
	IpamAllocationUpdateIPPoolConflictCounts    api.Int64Counter
	IpamAllocationErrInternalCounts             api.Int64Counter
	IpamAllocationErrNoAvailablePoolCounts      api.Int64Counter
	IpamAllocationErrRetriesExhaustedCounts     api.Int64Counter
	IpamAllocationErrIPUsedOutCounts            api.Int64Counter
	ipamAllocationAverageDurationSeconds        = new(asyncFloat64Gauge)
	ipamAllocationMaxDurationSeconds            = new(asyncFloat64Gauge)
	ipamAllocationMinDurationSeconds            = new(asyncFloat64Gauge)
	ipamAllocationLatestDurationSeconds         = new(asyncFloat64Gauge)
	ipamAllocationDurationSecondsHistogram      api.Float64Histogram
	ipamAllocationAverageLimitDurationSeconds   = new(asyncFloat64Gauge)
	ipamAllocationMaxLimitDurationSeconds       = new(asyncFloat64Gauge)
	ipamAllocationMinLimitDurationSeconds       = new(asyncFloat64Gauge)
	ipamAllocationLatestLimitDurationSeconds    = new(asyncFloat64Gauge)
	ipamAllocationLimitDurationSecondsHistogram api.Float64Histogram

	// ipam release metrics in spiderpool-agent
	IpamReleaseTotalCounts                   api.Int64Counter
	IpamReleaseFailureCounts                 api.Int64Counter
	IpamReleaseUpdateIPPoolConflictCounts    api.Int64Counter
	IpamReleaseErrInternalCounts             api.Int64Counter
	IpamReleaseErrRetriesExhaustedCounts     api.Int64Counter
	ipamReleaseAverageDurationSeconds        = new(asyncFloat64Gauge)
	ipamReleaseMaxDurationSeconds            = new(asyncFloat64Gauge)
	ipamReleaseMinDurationSeconds            = new(asyncFloat64Gauge)
	ipamReleaseLatestDurationSeconds         = new(asyncFloat64Gauge)
	ipamReleaseDurationSecondsHistogram      api.Float64Histogram
	ipamReleaseAverageLimitDurationSeconds   = new(asyncFloat64Gauge)
	ipamReleaseMaxLimitDurationSeconds       = new(asyncFloat64Gauge)
	ipamReleaseMinLimitDurationSeconds       = new(asyncFloat64Gauge)
	ipamReleaseLatestLimitDurationSeconds    = new(asyncFloat64Gauge)
	ipamReleaseLimitDurationSecondsHistogram api.Float64Histogram

	// IP GC metrics in spiderpool-controller
	IPGCTotalCounts   api.Int64Counter
	IPGCFailureCounts api.Int64Counter

	// IPPool&Subnet metrics in spiderpool-controller
	TotalIPPoolCounts       = new(asyncInt64Gauge)
	IPPoolTotalIPCounts     api.Int64Counter
	IPPoolAvailableIPCounts api.Int64Counter
	TotalSubnetCounts       = new(asyncInt64Gauge)
	SubnetPoolCounts        = new(asyncInt64Gauge)
	SubnetTotalIPCounts     api.Int64Counter
	SubnetAvailableIPCounts api.Int64Counter

	// SpiderSubnet feature performance monitoring metric in spiderpool-agent
	AutoPoolWaitedForAvailableCounts api.Int64Counter
)

// asyncFloat64Gauge is custom otel float64 gauge
type asyncFloat64Gauge struct {
	gaugeMetric           api.Float64ObservableGauge
	observerValueToReport float64
	observerAttrsToReport []attribute.KeyValue
	observerLock          lock.RWMutex
}

// initGauge will new an otel float64 gauge metric and register a call back function
func (a *asyncFloat64Gauge) initGauge(metricName string, description string, isDebugLevel bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Record uses otel async gauge observe function
func (a *asyncFloat64Gauge) Record(value float64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

// asyncInt64Gauge is custom otel int64 gauge
type asyncInt64Gauge struct {
	gaugeMetric           api.Int64ObservableGauge
	observerValueToReport int64
	observerAttrsToReport []attribute.KeyValue
	observerLock          lock.RWMutex
}

// initGauge will new an otel int64 gauge metric and register a call back function
func (a *asyncInt64Gauge) initGauge(metricName string, description string, isDebugLevel bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Record uses otel async gauge observe function
func (a *asyncInt64Gauge) Record(value int64, attrs ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

// InitSpiderpoolAgentMetrics serves for spiderpool agent metrics initialization
func InitSpiderpoolAgentMetrics(ctx context.Context, cache podownercache.CacheInterface) error {
	_ = "STUB: not implemented"
	// for rdma
	return nil
}

// InitSpiderpoolControllerMetrics serves for spiderpool-controller metrics initialization
func InitSpiderpoolControllerMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// initSpiderpoolAgentAllocationMetrics will init spiderpool-agent IPAM allocation metrics
func initSpiderpoolAgentAllocationMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	// spiderpool agent ipam allocation total counts, metric type "int64 counter"
	return nil
}

// spiderpool agent ipam allocation failure counts, metric type "int64 counter"

// spiderpool agent ipam allocation update IPPool conflict counts, metric type "int64 counter"

// spiderpool agent ipam allocation internal error counts, metric type "int64 counter"

// spiderpool agent ipam allocation no available IPPool error counts, metric type "int64 counter"

// spiderpool agent ipam allocation retries exhausted error counts, metric type "int64 counter"

// spiderpool agent ipam allocation IP addresses used out error counts, metric type "int64 counter"

// spiderpool agent ipam average allocation duration, metric type "float64 gauge"

// spiderpool agent ipam maximum allocation duration, metric type "float64 gauge"

// spiderpool agent ipam minimum allocation duration, metric type "float64 gauge"

// spiderpool agent ipam latest allocation duration, metric type "float64 gauge"

// spiderpool agent ipam allocation duration bucket, metric type "float64 histogram"

// spiderpool agent ipam average allocation limit duration, metric type "float64 gauge"

// spiderpool agent ipam maximum allocation limit duration, metric type "float64 gauge"

// spiderpool agent ipam minimum allocation limit duration, metric type "float64 gauge"

// spiderpool agent ipam latest allocation limit duration, metric type "float64 gauge"

// spiderpool agent ipam allocation limit duration bucket, metric type "float64 histogram"

// initSpiderpoolAgentReleaseMetrics will init spiderpool-agent IPAM release metrics
func initSpiderpoolAgentReleaseMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	// spiderpool agent ipam release total counts, metric type "int64 counter"
	return nil
}

// spiderpool agent ipam release failure counts, metric type "int64 counter"

// spiderpool agent ipam release update IPPool conflict counts, metric type "int64 counter"

// spiderpool agent ipam releasing internal error counts, metric type "int64 counter"

// spiderpool agent ipam releasing retries exhausted error counts, metric type "int64 counter"

// spiderpool agent ipam average release duration, metric type "float64 gauge"

// spiderpool agent ipam maximum release duration, metric type "float64 gauge"

// spiderpool agent ipam minimum allocation duration, metric type "float64 gauge"

// spiderpool agent ipam latest release duration, metric type "float64 gauge"

// spiderpool agent ipam allocation duration bucket, metric type "float64 histogram"

// spiderpool agent ipam average release limit duration, metric type "float64 gauge"

// spiderpool agent ipam maximum release limit duration, metric type "float64 gauge"

// spiderpool agent ipam minimum allocation limit duration, metric type "float64 gauge"

// spiderpool agent ipam latest release limit duration, metric type "float64 gauge"

// spiderpool agent ipam allocation limit duration bucket, metric type "float64 histogram"

// initSpiderpoolControllerGCMetrics will init spiderpool-controller IP gc metrics
func initSpiderpoolControllerGCMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func initSpiderpoolControllerCRMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
