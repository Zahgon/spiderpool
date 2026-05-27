// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	"context"
	// "fmt"
	"time"
)

type IPAMConfig struct {
	EnableIPv4 bool
	EnableIPv6 bool

	EnableSpiderSubnet                   bool
	EnableAutoPoolForApplication         bool
	EnableSpiderSubnetAutoPool           bool
	EnableStatefulSet                    bool
	EnableKubevirtStaticIP               bool
	EnableReleaseConflictIPsForStateless bool
	EnableIPConflictDetection            bool
	EnableGatewayDetection               bool

	OperationRetries     int
	OperationGapDuration time.Duration

	MultusClusterNetwork *string
	AgentNamespace       string
}

func setDefaultsForIPAMConfig(config IPAMConfig) IPAMConfig {
	_ = "STUB: not implemented"
	return *new(IPAMConfig)
}

func (c *IPAMConfig) checkIPVersionEnable(ctx context.Context, tt ToBeAllocateds) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *IPAMConfig) filterPoolMisspecified(ctx context.Context, t *ToBeAllocated) error {
	_ = "STUB: not implemented"
	return nil
}

// for dual stack environment, support only ipv4 or ipv6 address
// if c.EnableIPv4 && v4Count == 0 {
//	return fmt.Errorf("%w, IPv4 is enabled, but no IPv4 IPPool specified for NIC %s", constant.ErrWrongInput, t.NIC)
// }
// if c.EnableIPv6 && v6Count == 0 {
//	return fmt.Errorf("%w, IPv6 is enabled, but no IPv6 IPPool specified for NIC %s", constant.ErrWrongInput, t.NIC)
// }
