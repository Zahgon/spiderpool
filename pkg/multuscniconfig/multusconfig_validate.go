// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package multuscniconfig

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

var (
	cniTypeField         = field.NewPath("spec").Child("cniType")
	macvlanConfigField   = field.NewPath("spec").Child("macvlanConfig")
	ipvlanConfigField    = field.NewPath("spec").Child("ipvlanConfig")
	vlanConfigField      = field.NewPath("spec").Child("vlan")
	sriovConfigField     = field.NewPath("spec").Child("sriovConfig")
	ibsriovConfigField   = field.NewPath("spec").Child("ibsriovConfig")
	ipoibConfigField     = field.NewPath("spec").Child("ipoibConfig")
	ovsConfigField       = field.NewPath("spec").Child("ovsConfig")
	customCniConfigField = field.NewPath("spec").Child("customCniTypeConfig")
	chainCniConfigField  = field.NewPath("spec").Child("chainCNIJsonData")
	annotationField      = field.NewPath("metadata").Child("annotations")
)

func (mcw *MultusConfigWebhook) validate(ctx context.Context, oldMultusConfig, multusConfig *spiderpoolv2beta1.SpiderMultusConfig) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func checkExistedConfig(spec *spiderpoolv2beta1.MultusCNIConfigSpec, exclude string) bool {
	_ = "STUB: not implemented"
	return false
}

func validateCNIConfig(multusConfig *spiderpoolv2beta1.SpiderMultusConfig) *field.Error {
	_ = "STUB: not implemented"
	// with Kubernetes OpenAPI validation and Mutating Webhook, multusConfSpec.CniType must not be nil and default to "custom"
	return nil
}

// multusConfig.Spec.CustomCNIConfig can be empty

// verify that the data is a valid CNI format

func validateVlanCNIConfig(master []string, bond *spiderpoolv2beta1.BondConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateVlanID(vlanID int32) error { _ = "STUB: not implemented"; return nil }

func (mcw *MultusConfigWebhook) validateAnnotation(ctx context.Context, multusConfig *spiderpoolv2beta1.SpiderMultusConfig) *field.Error {
	_ = "STUB: not implemented"
	// Helper function to check net-attach-def existence and ownership
	return nil
}

// net-attach-def already exists and is managed by SpiderMultusConfig, do not allow the creation of SpiderMultusConfig to take over its management.

// The net-attach-def already exists and is not managed by SpiderMultusConfig, allow the creation of SpiderMultusConfig to take over its management.

// Validate the annotation 'multus.spidernet.io/cr-name' to customize the net-attach-def resource name.

// Validate the custom net-attach-def CNI version

func validateCustomAnnoNameShouldNotBeChangeable(oldMultusConfig, newMultusConfig *spiderpoolv2beta1.SpiderMultusConfig) *field.Error {
	_ = "STUB: not implemented"
	return nil
}
