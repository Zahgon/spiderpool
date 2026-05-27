// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package multuscniconfig

import (
	"context"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

func mutateSpiderMultusConfig(ctx context.Context, smc *spiderpoolv2beta1.SpiderMultusConfig) {
	_ = "STUB: not implemented"
	return
}

// In the SpiderMultusConfig resource first creation, if we don't set `Spec.CniType` field, we need to set it to `custom`.
// The kubernetes webhook is called before OpenAPI JSONSchema validation

// with custom CNI configuration, we don't need to add Coordinator configuration

// inject the labels

func setMacvlanDefaultConfig(macvlanConfig *spiderpoolv2beta1.SpiderMacvlanCniConfig) {
	_ = "STUB: not implemented"
	return
}

func setBondDefaultConfig(bond *spiderpoolv2beta1.BondConfig) *spiderpoolv2beta1.BondConfig {
	_ = "STUB: not implemented"
	return nil
}

func setIPVlanDefaultConfig(ipvlanConfig *spiderpoolv2beta1.SpiderIPvlanCniConfig) {
	_ = "STUB: not implemented"
	return
}

func setVlanDefaultConfig(vlanConfig *spiderpoolv2beta1.SpiderVlanCniConfig) {
	_ = "STUB: not implemented"
	return
}

func setSriovDefaultConfig(sriovConfig *spiderpoolv2beta1.SpiderSRIOVCniConfig) {
	_ = "STUB: not implemented"
	return
}

func setIBSriovDefaultConfig(ibsriovConfig *spiderpoolv2beta1.SpiderIBSriovCniConfig) {
	_ = "STUB: not implemented"
	return
}

func setIpoibDefaultConfig(config *spiderpoolv2beta1.SpiderIpoibCniConfig) {
	_ = "STUB: not implemented"
	return
}

func setOvsDefaultConfig(ovsConfig *spiderpoolv2beta1.SpiderOvsCniConfig) {
	_ = "STUB: not implemented"
	return
}

func setCoordinatorDefaultConfig(coordinator *spiderpoolv2beta1.CoordinatorSpec) *spiderpoolv2beta1.CoordinatorSpec {
	_ = "STUB: not implemented"
	return nil
}
