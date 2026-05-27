// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	"context"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

func (i *ipam) Release(ctx context.Context, delArgs *models.IpamDelArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// If Pod still exists, change the timeout of ctx to be consistent with
// the deletion grace period of Pod. After this time, all IP allocation
// recycling should be completed by GC instead of CNI DEL.
//
// But if Pod no longer exists, CNI DEL is still called (CNI DEL may be
// called multiple times according to the CNI Specification), continue
// to use the original ctx of OAI UNIX client (default 30s).

// Give priority to the UID of Pod, and then consider ENV K8S_POD_UID in
// CNI_ARGS, because some CRIs do not set K8S_POD_UID (such as dockershim).
// If do not get UID through all the above channels, skip CNI DEL and hand
// over the task of IP allocation recycling to GC.

// if the kubevirt vm pod is not exist, the gc will release the legacy IP later

func (i *ipam) releaseForAllNICs(ctx context.Context, uid, nic string, endpoint *spiderpoolv2beta1.SpiderEndpoint) error {
	_ = "STUB: not implemented"
	return nil
}

// Check whether an StatefulSet needs to release its currently allocated IP addresses.
// It is discussed in https://github.com/spidernet-io/spiderpool/issues/1045

// Check whether the kubevirt VM pod needs to keep its IP allocation.

func (i *ipam) release(ctx context.Context, uid string, details []spiderpoolv2beta1.IPAllocationDetail) error {
	_ = "STUB: not implemented"
	return nil
}

// Record the metric of queuing time for release.

// ReleaseIPs will release the given IP corresponding NIC whole IPs,
// and we will release the SpiderEndpoint recorded IPs first and release the SpiderIPPool recorded IPs later.
func (i *ipam) ReleaseIPs(ctx context.Context, delArgs *models.IpamBatchDelArgs) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// we need to have the Pod UID for IP release operation

// set Pod UID to parameter

// check for release conflict IPs

// do not release conflict IPs for stateful Pod

// return error for 'IsReleaseConflictIPs'

// return error for 'IsReleaseConflictIPs'

// release stateless workload SpiderEndpoint IPs

// release IPPool IPs
