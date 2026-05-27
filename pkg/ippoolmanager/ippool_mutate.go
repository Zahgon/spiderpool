// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ippoolmanager

import (
	"context"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

func (iw *IPPoolWebhook) mutateIPPool(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) error {
	_ = "STUB: not implemented"
	return nil
}

// inherit gateway,vlan,routes from corresponding SpiderSubnet if not set

func (iw *IPPoolWebhook) setControllerSubnet(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) (*spiderpoolv2beta1.SpiderSubnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InheritSubnetProperties(subnet *spiderpoolv2beta1.SpiderSubnet, ipPool *spiderpoolv2beta1.SpiderIPPool) {
	_ = "STUB: not implemented"
	return
}

// if customer set empty route for this IPPool, it would not inherit the SpiderSubnet.Spec.Routes
