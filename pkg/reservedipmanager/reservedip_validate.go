// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package reservedipmanager

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

var (
	ipVersionField *field.Path = field.NewPath("spec").Child("ipVersion")
	ipsField       *field.Path = field.NewPath("spec").Child("ips")
)

func (rw *ReservedIPWebhook) validateCreateReservedIP(ctx context.Context, rIP *spiderpoolv2beta1.SpiderReservedIP) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func (rw *ReservedIPWebhook) validateUpdateReservedIP(ctx context.Context, oldRIP, newRIP *spiderpoolv2beta1.SpiderReservedIP) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateReservedIPShouldNotBeChanged(oldRIP, newRIP *spiderpoolv2beta1.SpiderReservedIP) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (rw *ReservedIPWebhook) validateReservedIPSpec(ctx context.Context, rIP *spiderpoolv2beta1.SpiderReservedIP) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (rw *ReservedIPWebhook) validateReservedIPIPVersion(version *types.IPVersion) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (rw *ReservedIPWebhook) validateReservedIPs(ctx context.Context, version types.IPVersion, ips []string) *field.Error {
	_ = "STUB: not implemented"
	return nil
}
