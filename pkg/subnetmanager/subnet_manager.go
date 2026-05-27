// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package subnetmanager

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/reservedipmanager"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

type SubnetManager interface {
	GetSubnetByName(ctx context.Context, subnetName string, cached bool) (*spiderpoolv2beta1.SpiderSubnet, error)
	ListSubnets(ctx context.Context, cached bool, opts ...client.ListOption) (*spiderpoolv2beta1.SpiderSubnetList, error)
	ReconcileAutoIPPool(ctx context.Context, pool *spiderpoolv2beta1.SpiderIPPool, subnetName string, podController types.PodTopController, autoPoolProperty types.AutoPoolProperty) (*spiderpoolv2beta1.SpiderIPPool, error)
}

type subnetManager struct {
	client     client.Client
	apiReader  client.Reader
	rIPManager reservedipmanager.ReservedIPManager
}

func NewSubnetManager(client client.Client, apiReader client.Reader, rIPManager reservedipmanager.ReservedIPManager) (SubnetManager, error) {
	_ = "STUB: not implemented"
	return *new(SubnetManager), nil
}

func (sm *subnetManager) GetSubnetByName(ctx context.Context, subnetName string, cached bool) (*spiderpoolv2beta1.SpiderSubnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sm *subnetManager) ListSubnets(ctx context.Context, cached bool, opts ...client.ListOption) (*spiderpoolv2beta1.SpiderSubnetList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sm *subnetManager) ReconcileAutoIPPool(ctx context.Context, pool *spiderpoolv2beta1.SpiderIPPool, subnetName string,
	podController types.PodTopController, autoPoolProperty types.AutoPoolProperty,
) (*spiderpoolv2beta1.SpiderIPPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if the pool needs to be created

// check if the given pool's IPs numbers are equal with the desired IP number counts

// refresh the label "ipam.spidernet.io/owner-application-uid" and "ipam.spidernet.io/ippool-reclaim"

// Vlan:        subnet.Spec.Vlan,

// label IPPoolCIDR

// set owner reference

// set finalizer

// give the auto-created IPPool a symbol to explain whether it is IP number flexible or fixed.

// preAllocateIPsFromSubnet will calculate the auto-created IPPool required IPs from corresponding SpiderSubnet and return it.
func (sm *subnetManager) preAllocateIPsFromSubnet(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet, pool *spiderpoolv2beta1.SpiderIPPool, ipVersion types.IPVersion, desiredIPNum int, podController types.PodTopController) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In the last reconcile process, the SpiderSubnet allocated IPs successfully but the pool creation process failed.

// If we have difference sets, which means the subnet updated its status successfully in the last shrink operation but the next ippool update operation failed.
// In the situation, the ippool may allocate or release one of ips that subnet updated. So, we should correct the subnet status.

// shrink: free IP number >= return IP Num
// when it needs to scale down IP, enough IP is released to make sure it scale down successfully

// exist auto pool allocated IPs

// free IPs

// filter reserved IPs

// check the filtered subnet free IP number is enough or not

func subnetStatusCount(subnet *spiderpoolv2beta1.SpiderSubnet) (totalCount, allocatedCount int64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// allocated IP Count
