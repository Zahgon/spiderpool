// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package reservedipmanager

import (
	"context"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

func (rw *ReservedIPWebhook) mutateReservedIP(ctx context.Context, rIP *spiderpoolv2beta1.SpiderReservedIP) error {
	_ = "STUB: not implemented"
	return nil
}
