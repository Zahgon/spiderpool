// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

func (i *ipam) getPoolCandidates(ctx context.Context, addArgs *models.IpamAddArgs, pod *corev1.Pod, podController types.PodTopController) (ToBeAllocateds, error) {
	_ = "STUB: not implemented"
	// If feature SpiderSubnet is enabled, select IPPool candidates through the
	// Pod annotations "ipam.spidernet.io/subnet" or "ipam.spidernet.io/subnets". (expect orphan Pod controller)
	return *new(ToBeAllocateds), nil
}

// Select IPPool candidates through the Pod annotation "ipam.spidernet.io/ippools".

// Select IPPool candidates through the Pod annotation "ipam.spidernet.io/ippool".

// Select IPPool candidates through the Namespace annotations
// "ipam.spidernet.io/default-ipv4-ippool" and "ipam.spidernet.io/default-ipv6-ippool".

// Select IPPool candidates through CNI network configuration.

// Select IPPools whose spec.default is true.

func (i *ipam) getPoolFromSubnetAnno(ctx context.Context, pod *corev1.Pod, nic string, cleanGateway bool, podController types.PodTopController) (*ToBeAllocated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get SpiderSubnet configuration from pod annotation

// default IPPool mode

// The SpiderSubnet feature doesn't support orphan Pod.

// This only serves for third party controller application, because we'll create or scale the auto-created IPPool here.
// For those kubernetes applications(such as deployment and replicaset), the spiderpool-controller will create or scale the auto-created IPPool asynchronously.

// if enableIPv4 is off and get the specified SpiderSubnet IPv4 name, just filter it out

// if enableIPv6 is off and get the specified SpiderSubnet IPv6 name, just filter it out

// findAppAutoPool only fetches kubernetes basic controller(like Deployment, StatefulSet etc...) corresponding auto-created IPPools.
func (i *ipam) findAppAutoPool(ctx context.Context, subnetName, ifName, labelIPPoolIPVersionValue string, desiredIPNumber int, podController types.PodTopController) (*spiderpoolv2beta1.SpiderIPPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we fetched Auto-created IPPool but it doesn't have any IPs, just wait for a while and let the IPPool informer to allocate IPs for it

// applyThirdControllerAutoPool will fetch or reconcile third-party controller corresponding auto-created IPPools,
// and the kubernetes basic controller like Deployment,StatefulSet etc... We'll reconcile their auto-created IPPools in spiderpool-controller component.
func (i *ipam) applyThirdControllerAutoPool(ctx context.Context, subnetName string, podController types.PodTopController, autoPoolProperty types.AutoPoolProperty) (*spiderpoolv2beta1.SpiderIPPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ipam) getPoolFromPodAnnoPools(ctx context.Context, anno, currentNIC string) (ToBeAllocateds, error) {
	_ = "STUB: not implemented"
	return *new(ToBeAllocateds), nil
}

// validate and mutate the IPPools annotation value

func (i *ipam) getPoolFromPodAnnoPool(ctx context.Context, anno, nic string, cleanGateway bool) (*ToBeAllocated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check IPv4 PoolName wildcard

// overwrite the annoPodIPPool

// check IPv6 PoolName wildcard

// overwrite the annoPodIPPool

func (i *ipam) getPoolFromNS(ctx context.Context, namespace, nic string, cleanGateway bool) (*ToBeAllocated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ipam) getPoolFromNetConf(ctx context.Context, nic string, netConfV4Pool, netConfV6Pool []string, cleanGateway bool) (*ToBeAllocated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ipam) getClusterDefaultPools(ctx context.Context, nic string, cleanGateway bool) (*ToBeAllocated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
