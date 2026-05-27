// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	"context"
	"net"

	. "github.com/onsi/gomega"
	"github.com/spidernet-io/spiderpool/pkg/lock"

	spiderpool "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"

	. "github.com/onsi/ginkgo/v2"
	frame "github.com/spidernet-io/e2eframework/framework"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var usedSubnetsLock = new(lock.Mutex)

func GenerateExampleV4SubnetObject(f *frame.Framework, ipNum int) (string, *spiderpool.SpiderSubnet) {
	_ = "STUB: not implemented"
	return "", nil
}

func GenerateExampleV6SubnetObject(f *frame.Framework, ipNum int) (string, *spiderpool.SpiderSubnet) {
	_ = "STUB: not implemented"
	return "", nil
}

func CreateSubnet(f *frame.Framework, subnet *spiderpool.SpiderSubnet, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to wait for finish last deleting

func WaitCreateSubnetUntilFinish(ctx context.Context, f *frame.Framework, subnet *spiderpool.SpiderSubnet, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteSubnetByName(f *frame.Framework, subnetName string, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func GetSubnetByName(f *frame.Framework, subnetName string) (*spiderpool.SpiderSubnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteSubnetUntilFinish(ctx context.Context, f *frame.Framework, subnetName string, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitValidateSubnetAllocatedIPCount(ctx context.Context, f *frame.Framework, subnetName string, allocatedIPCount int64) error {
	_ = "STUB: not implemented"
	return nil
}

// The informer of SpiderSubnet will delay synchronizing its own state information
// which may cause failure 'runtime error: invalid memory address or nil pointer dereference'

func PatchSpiderSubnet(f *frame.Framework, desiredSubnet, originalSubnet *spiderpool.SpiderSubnet, opts ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitIppoolNumberInSubnet(ctx context.Context, f *frame.Framework, subnetName string, poolNums int) error {
	_ = "STUB: not implemented"
	return nil
}

func GetAvailableIpsInSubnet(f *frame.Framework, subnetName string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WaitValidateSubnetAndPoolIPConsistency(ctx context.Context, f *frame.Framework, subnetName string) error {
	_ = "STUB: not implemented"
	return nil
}

func BatchCreateSubnet(f *frame.Framework, version types.IPVersion, subnetNums, subnetIPNums int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetAllSubnet(f *frame.Framework, opts ...client.ListOption) (*spiderpool.SpiderSubnetList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
