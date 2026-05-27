// Copyright 2024 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package ethtool

import (
	"github.com/spidernet-io/spiderpool/pkg/rdmametrics/oteltype"
)

func Stats(netIfName string) ([]oteltype.Metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// speed unknown = 4294967295

func expectPriorityMetrics(s string) bool { _ = "STUB: not implemented"; return false }

func extractNameWithPriority(str string) (name string, priority int, ok bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}
