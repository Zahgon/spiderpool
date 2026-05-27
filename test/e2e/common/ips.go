// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Package common provides test utilities for E2E tests.
package common

import (
	"net"

	"github.com/spidernet-io/spiderpool/pkg/lock"
)

var generatedIPs map[string]bool

var generateIPsLock = new(lock.Mutex)

func GenerateIPs(cidr string, num int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Increment the network IP to start from 192.168.0.1 instead of 192.168.0.0

func ipCount(network *net.IPNet) int { _ = "STUB: not implemented"; return 0 }

func inc(ip net.IP) { _ = "STUB: not implemented"; return }
