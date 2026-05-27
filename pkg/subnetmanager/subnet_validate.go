// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package subnetmanager

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

var (
	ipVersionField         *field.Path = field.NewPath("spec").Child("ipVersion")
	subnetField            *field.Path = field.NewPath("spec").Child("subnet")
	ipsField               *field.Path = field.NewPath("spec").Child("ips")
	excludeIPsField        *field.Path = field.NewPath("spec").Child("excludeIPs")
	gatewayField           *field.Path = field.NewPath("spec").Child("gateway")
	routesField            *field.Path = field.NewPath("spec").Child("routes")
	controlledIPPoolsField *field.Path = field.NewPath("status").Child("controlledIPPools")
)

func (sw *SubnetWebhook) validateCreateSubnet(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func (sw *SubnetWebhook) validateUpdateSubnet(ctx context.Context, oldSubnet, newSubnet *spiderpoolv2beta1.SpiderSubnet) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateSubnetShouldNotBeChanged(oldSubnet, newSubnet *spiderpoolv2beta1.SpiderSubnet) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *SubnetWebhook) validateSubnetSpec(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubnetIPInUse(subnet *spiderpoolv2beta1.SpiderSubnet) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *SubnetWebhook) validateSubnetIPVersion(version *types.IPVersion) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *SubnetWebhook) validateSubnetCIDR(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

// since we met already exist Subnet resource, we just return the error to avoid the following taxing operations.
// the user can also use k8s 'errors.IsAlreadyExists' to get the right error type assertion.

// validateOrphanIPPool will check the SpiderSubnet.Spec.Subnet whether overlaps with the cluster orphan SpiderIPPool.Spec.Subnet.
// And we also require the IPPool.Spec.IPs belong to Subnet.Spec.IPs if they are in the same subnet
func (sw *SubnetWebhook) validateOrphanIPPool(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

// validate the Spec.Subnet whether overlaps or not

// validate the Spec.IPs whether contains or not

func validateSubnetIPs(version types.IPVersion, subnet string, ips []string) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubnetExcludeIPs(version types.IPVersion, subnet string, excludeIPs []string) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubnetGateway(subnet *spiderpoolv2beta1.SpiderSubnet) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateSubnetRoutes(version types.IPVersion, subnet string, routes []spiderpoolv2beta1.Route) *field.Error {
	_ = "STUB: not implemented"
	return nil
}
