// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

type ToBeAllocateds []*ToBeAllocated

func (tt *ToBeAllocateds) Pools() []string { _ = "STUB: not implemented"; return nil }

func (tt *ToBeAllocateds) Candidates() []*PoolCandidate { _ = "STUB: not implemented"; return nil }

type ToBeAllocated struct {
	NIC            string
	CleanGateway   bool
	PoolCandidates []*PoolCandidate
}

func (t *ToBeAllocated) Pools() []string { _ = "STUB: not implemented"; return nil }

func (t *ToBeAllocated) String() string { _ = "STUB: not implemented"; return "" }

type PoolCandidate struct {
	IPVersion types.IPVersion
	Pools     []string
	PToIPPool PoolNameToIPPool
}

func (c *PoolCandidate) String() string { _ = "STUB: not implemented"; return "" }

type PoolNameToIPPool map[string]*spiderpoolv2beta1.SpiderIPPool

func (pp *PoolNameToIPPool) IPPools() []*spiderpoolv2beta1.SpiderIPPool {
	_ = "STUB: not implemented"
	return nil
}

func (pp PoolNameToIPPool) String() string { _ = "STUB: not implemented"; return "" }
