// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package convert

import (
	"net"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

func ConvertIPDetailsToIPConfigsAndAllRoutes(details []spiderpoolv2beta1.IPAllocationDetail, enableIPConflictDetection, enableGatewayDetection bool) ([]*models.IPConfig, []*models.Route) {
	_ = "STUB: not implemented"
	return nil, nil
}

// don't set default route if we have cleangateway

// don't set default route if we have cleangateway

func ConvertResultsToIPConfigsAndAllRoutes(results []*types.AllocationResult) ([]*models.IPConfig, []*models.Route) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genDefaultRoute(nic, gateway string) *models.Route { _ = "STUB: not implemented"; return nil }

func ConvertResultsToIPDetails(results []*types.AllocationResult, isMultipleNicWithNoName bool) []spiderpoolv2beta1.IPAllocationDetail {
	_ = "STUB: not implemented"
	return nil
}

// Set MAC if provided and not already set

// Extract MAC if provided

// If no NIC name, sort the results with NIC and specify the first NIC name with "eth0", the others with empty.
// For the other NIC allocation, we'll retrieve the results from the Endpoint resource and update the Endpoint resource with real NIC name for the no NIC name set.

func ConvertAnnoPodRoutesToOAIRoutes(annoPodRoutes types.AnnoPodRoutesValue) []*models.Route {
	_ = "STUB: not implemented"
	return nil
}

func ConvertSpecRoutesToOAIRoutes(nic string, specRoutes []spiderpoolv2beta1.Route) []*models.Route {
	_ = "STUB: not implemented"
	return nil
}

func ConvertOAIRoutesToSpecRoutes(oaiRoutes []*models.Route) []spiderpoolv2beta1.Route {
	_ = "STUB: not implemented"
	return nil
}

func GroupIPAllocationDetails(uid string, details []spiderpoolv2beta1.IPAllocationDetail) types.PoolNameToIPAndUIDs {
	_ = "STUB: not implemented"
	return *new(types.PoolNameToIPAndUIDs)
}

func GenIPConfigResult(allocateIP net.IP, nic string, ipPool *spiderpoolv2beta1.SpiderIPPool) *models.IPConfig {
	_ = "STUB: not implemented"
	return nil
}

func UnmarshalIPPoolAllocatedIPs(data *string) (spiderpoolv2beta1.PoolIPAllocations, error) {
	_ = "STUB: not implemented"
	return *new(spiderpoolv2beta1.PoolIPAllocations), nil
}

func MarshalIPPoolAllocatedIPs(records spiderpoolv2beta1.PoolIPAllocations) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalSubnetAllocatedIPPools(data *string) (spiderpoolv2beta1.PoolIPPreAllocations, error) {
	_ = "STUB: not implemented"
	return *new(spiderpoolv2beta1.PoolIPPreAllocations), nil
}

func MarshalSubnetAllocatedIPPools(preAllocations spiderpoolv2beta1.PoolIPPreAllocations) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
