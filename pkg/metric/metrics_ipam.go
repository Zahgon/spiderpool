// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package metric

import (
	"context"

	"github.com/spidernet-io/spiderpool/pkg/lock"
)

// IPAMDurationConstruct is Singleton
var IPAMDurationConstruct = new(ipamDurationConstruct)

type ipamDurationConstruct struct {
	allocate      durationConstruct
	release       durationConstruct
	allocateLimit durationConstruct
	releaseLimit  durationConstruct
}

type durationConstruct struct {
	cacheLock lock.RWMutex

	avgDuration float64
	maxDuration float64
	minDuration float64

	counts int
}

// RecordIPAMAllocationDuration serves for spiderpool agent IPAM allocation.
func (idc *ipamDurationConstruct) RecordIPAMAllocationDuration(ctx context.Context, allocationDuration float64) {
	_ = "STUB: not implemented"
	return
}

// latest allocation duration

// allocation duration histogram

// IPAM average allocation duration

// IPAM maximum allocation duration

// IPAM minimum allocation duration

// RecordIPAMReleaseDuration serves for spiderpool agent IPAM allocation.
func (idc *ipamDurationConstruct) RecordIPAMReleaseDuration(ctx context.Context, releaseDuration float64) {
	_ = "STUB: not implemented"
	return
}

// latest release duration

// release duration histogram

// IPAM average release duration

// IPAM maximum release duration

// IPAM minimum release duration

func (idc *ipamDurationConstruct) RecordIPAMAllocationLimitDuration(ctx context.Context, limitDuration float64) {
	_ = "STUB: not implemented"
	return
}

func (idc *ipamDurationConstruct) RecordIPAMReleaseLimitDuration(ctx context.Context, limitDuration float64) {
	_ = "STUB: not implemented"
	return
}
