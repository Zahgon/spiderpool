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

func getCustomRoutes(pod *corev1.Pod) ([]*models.Route, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func groupCustomRoutes(ctx context.Context, customRoutes []*models.Route, results []*types.AllocationResult) error {
	_ = "STUB: not implemented"
	return nil
}

// getAutoPoolIPNumber calculates the auto-created IPPool IP number with the given params pod and pod top controller.
// If it's an orphan pod, it will return 1.
func getAutoPoolIPNumber(pod *corev1.Pod, podController types.PodTopController) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// orphan pod

// check out negative number

// fixed IP number, just return it

// flexible IP Number

// third party controller only supports fixed auto-created IPPool IP number

// collect application replicas and custom flexible IP number

// isPoolIPsDesired checks the auto-created IPPool's IPs whether matches its AutoDesiredIPCount
func isPoolIPsDesired(pool *spiderpoolv2beta1.SpiderIPPool, desiredIPCount int) bool {
	_ = "STUB: not implemented"
	return false
}

func IsMultipleNicWithNoName(anno map[string]string) bool { _ = "STUB: not implemented"; return false }

func validateAndMutateMultipleNICAnnotations(annoIPPoolsValue types.AnnoPodIPPoolsValue, currentNIC string) error {
	_ = "STUB: not implemented"
	return nil
}

// require all item NIC should be same in specified or unspecified.

// once we met the unspecified NIC in multiple NIC with no name specified mode, we give it an index name.
// since the latter allocation will be executed in concurrency for all NICs in the first cmdAdd,
// we could sort the allocation results in sequence by the NIC name.

// require no NIC name duplicated, this works for multiple NIC specified by the users
