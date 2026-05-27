// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package subnetmanager

import (
	"context"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

func (sw *SubnetWebhook) mutateSubnet(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) error {
	_ = "STUB: not implemented"
	return nil
}
