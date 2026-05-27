// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package types

import "github.com/spidernet-io/spiderpool/api/v1/agent/models"

type IPVersion = int64

type Vlan = int64

type AllocationResult struct {
	IP           *models.IPConfig
	Routes       []*models.Route
	CleanGateway bool
}

type IPAndUID struct {
	IP  string
	UID string
}

type PoolNameToIPAndUIDs map[string][]IPAndUID

func (pius *PoolNameToIPAndUIDs) Pools() []string { _ = "STUB: not implemented"; return nil }
