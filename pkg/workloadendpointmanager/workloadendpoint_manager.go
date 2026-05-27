// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package workloadendpointmanager

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

type WorkloadEndpointManager interface {
	GetEndpointByName(ctx context.Context, namespace, podName string, cached bool) (*spiderpoolv2beta1.SpiderEndpoint, error)
	ListEndpoints(ctx context.Context, cached bool, opts ...client.ListOption) (*spiderpoolv2beta1.SpiderEndpointList, error)
	DeleteEndpoint(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint) error
	RemoveFinalizer(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint) error
	PatchIPAllocationResults(ctx context.Context, results []*types.AllocationResult, endpoint *spiderpoolv2beta1.SpiderEndpoint, pod *corev1.Pod, podController types.PodTopController, isMultipleNicWithNoName bool) error
	ReallocateCurrentIPAllocation(ctx context.Context, uid, nodeName, nic string, endpoint *spiderpoolv2beta1.SpiderEndpoint, isMultipleNicWithNoName bool) error
	UpdateAllocationNICName(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint, nic string) (*spiderpoolv2beta1.PodIPAllocation, error)
	ReleaseEndpointIPs(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint, uid string) ([]spiderpoolv2beta1.IPAllocationDetail, error)
	ReleaseEndpointAndFinalizer(ctx context.Context, namespace, podName string, cached bool) error
	PatchEndpointAllocationIPs(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint, endpointIPs []spiderpoolv2beta1.IPAllocationDetail) error
}

type workloadEndpointManager struct {
	client    client.Client
	apiReader client.Reader

	enableStatefulSet      bool
	enableKubevirtStaticIP bool
}

func NewWorkloadEndpointManager(client client.Client, apiReader client.Reader, enableStatefulSet, enableKubevirtStaticIP bool) (WorkloadEndpointManager, error) {
	_ = "STUB: not implemented"
	return *new(WorkloadEndpointManager), nil
}

func (em *workloadEndpointManager) GetEndpointByName(ctx context.Context, namespace, podName string, cached bool) (*spiderpoolv2beta1.SpiderEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (em *workloadEndpointManager) ListEndpoints(ctx context.Context, cached bool, opts ...client.ListOption) (*spiderpoolv2beta1.SpiderEndpointList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (em *workloadEndpointManager) DeleteEndpoint(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint) error {
	_ = "STUB: not implemented"
	return nil
}

func (em *workloadEndpointManager) RemoveFinalizer(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint) error {
	_ = "STUB: not implemented"
	return nil
}

func (em *workloadEndpointManager) PatchIPAllocationResults(ctx context.Context, results []*types.AllocationResult, endpoint *spiderpoolv2beta1.SpiderEndpoint, pod *corev1.Pod, podController types.PodTopController, isMultipleNicWithNoName bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not set ownerReference for Endpoint when its corresponding Pod is
// controlled by StatefulSet/KubevirtVMI. Once the Pod of StatefulSet/KubevirtVMI is recreated,
// we can immediately retrieve the old IP allocation results from the
// Endpoint without worrying about the cascading deletion of the Endpoint.

// Using ipam.spidernet.io/ippools to specify multiple NICs,
// if only one NIC's IP pool changes, only the changed NIC needs to have its IP address reassigned, while the other NICs remain unaffected.

func (em *workloadEndpointManager) ReallocateCurrentIPAllocation(ctx context.Context, uid, nodeName, nic string, endpoint *spiderpoolv2beta1.SpiderEndpoint, isMultipleNicWithNoName bool) error {
	_ = "STUB: not implemented"
	return nil
}

// For the multiple NICs allocation with no NIC name specified,
// we'll allocate all NICs IPs in the first allocation and record the details in the Endpoint resource.
// This is NIC must be "eth0" and others with empty NIC name, we reset the real NIC name for the empties in the retrieve actions.

func (em *workloadEndpointManager) UpdateAllocationNICName(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint, nic string) (*spiderpoolv2beta1.PodIPAllocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the SpiderEndpoint status allocation is already in order by NIC sequence

// PatchEndpointAllocationIPs will patch the SpiderEndpoint status recorded IPs.
func (em *workloadEndpointManager) PatchEndpointAllocationIPs(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint, newEndpointIPs []spiderpoolv2beta1.IPAllocationDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// ReleaseEndpointIPs will release the SpiderEndpoint status recorded IPs.
func (em *workloadEndpointManager) ReleaseEndpointIPs(ctx context.Context, endpoint *spiderpoolv2beta1.SpiderEndpoint, podUID string) ([]spiderpoolv2beta1.IPAllocationDetail, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (em *workloadEndpointManager) ReleaseEndpointAndFinalizer(ctx context.Context, namespace, podName string, cached bool) error {
	_ = "STUB: not implemented"
	return nil
}
