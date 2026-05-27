// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	"context"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
	"github.com/spidernet-io/spiderpool/pkg/ippoolmanager"
	"github.com/spidernet-io/spiderpool/pkg/kubevirtmanager"
	"github.com/spidernet-io/spiderpool/pkg/limiter"
	"github.com/spidernet-io/spiderpool/pkg/lock"
	"github.com/spidernet-io/spiderpool/pkg/namespacemanager"
	"github.com/spidernet-io/spiderpool/pkg/nodemanager"
	"github.com/spidernet-io/spiderpool/pkg/podmanager"
	"github.com/spidernet-io/spiderpool/pkg/statefulsetmanager"
	"github.com/spidernet-io/spiderpool/pkg/subnetmanager"
	"github.com/spidernet-io/spiderpool/pkg/types"
	"github.com/spidernet-io/spiderpool/pkg/workloadendpointmanager"
)

type IPAM interface {
	Allocate(ctx context.Context, addArgs *models.IpamAddArgs) (*models.IpamAddResponse, error)
	Release(ctx context.Context, delArgs *models.IpamDelArgs) error
	ReleaseIPs(ctx context.Context, delArgs *models.IpamBatchDelArgs) error
	Start(ctx context.Context) error
}

type ipam struct {
	config      IPAMConfig
	ipamLimiter limiter.Limiter
	failure     *failureCache

	ipPoolManager   ippoolmanager.IPPoolManager
	endpointManager workloadendpointmanager.WorkloadEndpointManager
	nodeManager     nodemanager.NodeManager
	nsManager       namespacemanager.NamespaceManager
	podManager      podmanager.PodManager
	stsManager      statefulsetmanager.StatefulSetManager
	subnetManager   subnetmanager.SubnetManager
	kubevirtManager kubevirtmanager.KubevirtManager
}

func NewIPAM(
	config IPAMConfig,
	ipPoolManager ippoolmanager.IPPoolManager,
	endpointManager workloadendpointmanager.WorkloadEndpointManager,
	nodeManager nodemanager.NodeManager,
	nsManager namespacemanager.NamespaceManager,
	podManager podmanager.PodManager,
	stsManager statefulsetmanager.StatefulSetManager,
	subnetManager subnetmanager.SubnetManager,
	kubevirtManager kubevirtmanager.KubevirtManager,
) (IPAM, error) {
	_ = "STUB: not implemented"
	return *new(IPAM), nil
}

func (i *ipam) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type failureCache struct {
	l       lock.RWMutex
	entries map[string][]*types.AllocationResult
}

func newFailureCache() *failureCache { _ = "STUB: not implemented"; return nil }

func (c *failureCache) addFailureIPs(uid string, results []*types.AllocationResult) {
	_ = "STUB: not implemented"
	return
}

func (c *failureCache) rmFailureIPs(uid string) { _ = "STUB: not implemented"; return }

func (c *failureCache) getFailureIPs(uid string) []*types.AllocationResult {
	_ = "STUB: not implemented"
	return nil
}
