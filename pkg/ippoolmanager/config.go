// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ippoolmanager

const (
	defaultMaxAllocatedIPs = 5000
)

type IPPoolManagerConfig struct {
	MaxAllocatedIPs           *int
	EnableKubevirtStaticIP    bool
	EnableGatewayDetection    bool
	EnableIPConflictDetection bool
}

func setDefaultsForIPPoolManagerConfig(config IPPoolManagerConfig) IPPoolManagerConfig {
	_ = "STUB: not implemented"
	return *new(IPPoolManagerConfig)
}
