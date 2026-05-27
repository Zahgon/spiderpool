// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ippoolmanager

import (
	"sort"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

func IsAutoCreatedIPPool(pool *spiderpoolv2beta1.SpiderIPPool) bool {
	_ = "STUB: not implemented"
	// only the auto-created IPPool owns the annotation "ipam.spidernet.io/owner-application"
	return false
}

func NewAutoPoolPodAffinity(podTopController types.PodTopController) *metav1.LabelSelector {
	_ = "STUB: not implemented"
	return nil
}

func IsMatchAutoPoolAffinity(podAffinity *metav1.LabelSelector, podTopController types.PodTopController) bool {
	_ = "STUB: not implemented"
	return false
}

// ByPoolPriority implements sort.Interface
var _ sort.Interface = &ByPoolPriority{}

type ByPoolPriority []*spiderpoolv2beta1.SpiderIPPool

func (b ByPoolPriority) Len() int { _ = "STUB: not implemented"; return 0 }

func (b ByPoolPriority) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (b ByPoolPriority) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// Pod Affinity
	return false
}

// Node Affinity

// Namespace Affinity

// Multus Name

// findAllocatedIPFromRecords try to find pod NIC previous allocated IP from the IPPool.Status.AllocatedIPs
// this function serves for the issue: https://github.com/spidernet-io/spiderpool/issues/2517
func findAllocatedIPFromRecords(allocatedRecords spiderpoolv2beta1.PoolIPAllocations, namespacedName, podUID string) (previousIP string, hasFound bool) {
	_ = "STUB: not implemented"
	return "", false
}

// HasWildcardInStr checks whether the wildcard '*', '?', '[]' exists in the given string variable
func HasWildcardInStr(str string) bool { _ = "STUB: not implemented"; return false }

func HasWildcardInSlice(arr []string) bool { _ = "STUB: not implemented"; return false }
