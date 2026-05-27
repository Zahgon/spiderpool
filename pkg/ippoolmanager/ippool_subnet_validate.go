// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ippoolmanager

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

func (iw *IPPoolWebhook) validateCreateIPPoolWhileEnableSpiderSubnet(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func (iw *IPPoolWebhook) validateUpdateIPPoolWhileEnableSpiderSubnet(ctx context.Context, oldIPPool, newIPPool *spiderpoolv2beta1.SpiderIPPool) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func (iw *IPPoolWebhook) validateSubnetTotalIPsContainsIPPoolTotalIPs(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateNewAutoPoolTotalIPsWithinSubnet(pool *spiderpoolv2beta1.SpiderIPPool, subnet *spiderpoolv2beta1.SpiderSubnet) *field.Error {
	_ = "STUB: not implemented"
	return nil
}
