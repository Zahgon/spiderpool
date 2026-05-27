// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ippoolmanager

import (
	"context"
	"net"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/reservedipmanager"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

type IPPoolManager interface {
	GetIPPoolByName(ctx context.Context, poolName string, cached bool) (*spiderpoolv2beta1.SpiderIPPool, error)
	ListIPPools(ctx context.Context, cached bool, opts ...client.ListOption) (*spiderpoolv2beta1.SpiderIPPoolList, error)
	AllocateIP(ctx context.Context, poolName, nic string, pod *corev1.Pod, podController types.PodTopController) (*models.IPConfig, error)
	ReleaseIP(ctx context.Context, poolName string, ipAndUIDs []types.IPAndUID) error
	UpdateAllocatedIPs(ctx context.Context, poolName, namespacedName string, ipAndCIDs []types.IPAndUID) error
	ParseWildcardPoolNameList(ctx context.Context, PoolNames []string, ipVersion types.IPVersion) (newPoolNames []string, hasWildcard bool, err error)
}

type ipPoolManager struct {
	config     IPPoolManagerConfig
	client     client.Client
	apiReader  client.Reader
	rIPManager reservedipmanager.ReservedIPManager
}

func NewIPPoolManager(config IPPoolManagerConfig, client client.Client, apiReader client.Reader, rIPManager reservedipmanager.ReservedIPManager) (IPPoolManager, error) {
	_ = "STUB: not implemented"
	return *new(IPPoolManager), nil
}

func (im *ipPoolManager) GetIPPoolByName(ctx context.Context, poolName string, cached bool) (*spiderpoolv2beta1.SpiderIPPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (im *ipPoolManager) ListIPPools(ctx context.Context, cached bool, opts ...client.ListOption) (*spiderpoolv2beta1.SpiderIPPoolList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (im *ipPoolManager) AllocateIP(ctx context.Context, poolName, nic string, pod *corev1.Pod, podController types.PodTopController) (*models.IPConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(@cyclinder): set these values from ippool.spec

func (im *ipPoolManager) genRandomIP(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool, pod *corev1.Pod, podController types.PodTopController) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}

// In a multi-NIC scenario, if one of the NIC pools does not have enough IPs, an allocation failure message will be displayed.
// However, other IP pools still have IPs, which will cause IPs in other pools to be exhausted.
// Check if there is a duplicate Pod UID in IPPool.allocatedRecords.
// If so, we skip this allocation and assume that this Pod has already obtained an IP address in the pool.

// traverse the usedIPs to find the previous allocated IPs if there be
// reference issue: https://github.com/spidernet-io/spiderpool/issues/2517

// reference issue: https://github.com/spidernet-io/spiderpool/issues/3771

// Adding a newly assigned IP

func (im *ipPoolManager) ReleaseIP(ctx context.Context, poolName string, ipAndUIDs []types.IPAndUID) error {
	_ = "STUB: not implemented"
	return nil
}

// reference issue: https://github.com/spidernet-io/spiderpool/issues/3771

func (im *ipPoolManager) UpdateAllocatedIPs(ctx context.Context, poolName, namespacedName string, ipAndUIDs []types.IPAndUID) error {
	_ = "STUB: not implemented"
	return nil
}

func (im *ipPoolManager) ParseWildcardPoolNameList(ctx context.Context, poolNamesArr []string, ipVersion types.IPVersion) (newPoolNames []string, hasWildcard bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// wildcard matches

// original IPPool name
