// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ippoolmanager

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

var (
	ipVersionField   *field.Path = field.NewPath("spec").Child("ipVersion")
	subnetField      *field.Path = field.NewPath("spec").Child("subnet")
	ipsField         *field.Path = field.NewPath("spec").Child("ips")
	gatewayField     *field.Path = field.NewPath("spec").Child("gateway")
	routesField      *field.Path = field.NewPath("spec").Child("routes")
	podAffinityField *field.Path = field.NewPath("spec").Child("podAffinity")
)

func (iw *IPPoolWebhook) validateCreateIPPool(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func (iw *IPPoolWebhook) validateUpdateIPPool(ctx context.Context, oldIPPool, newIPPool *spiderpoolv2beta1.SpiderIPPool) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateIPPoolShouldNotBeChanged(oldIPPool, newIPPool *spiderpoolv2beta1.SpiderIPPool) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (iw *IPPoolWebhook) validateIPPoolSpec(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateIPPoolIPInUse(ipPool *spiderpoolv2beta1.SpiderIPPool) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (iw *IPPoolWebhook) validateIPPoolIPVersion(version *types.IPVersion) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (iw *IPPoolWebhook) validateIPPoolCIDR(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

// since we met already exist IPPool resource, we just return the error to avoid the following taxing operations.
// the user can also use k8s 'errors.IsAlreadyExists' to get the right error type assertion.

func (iw *IPPoolWebhook) validateIPPoolAvailableIPs(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateIPPoolGateway(ipPool *spiderpoolv2beta1.SpiderIPPool) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateIPPoolRoutes(version types.IPVersion, subnet string, routes []spiderpoolv2beta1.Route) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateContainsIPRange(fieldPath *field.Path, version types.IPVersion, subnet string, ipRange string) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateContainsIP(fieldPath *field.Path, version types.IPVersion, subnet string, ip string) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateIPPoolPodAffinity(fieldPath *field.Path, ipPool *spiderpoolv2beta1.SpiderIPPool) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// auto-created IPPool special podAffinity validation

// normal IPPool podAffinity validation
