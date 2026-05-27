// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package coordinatormanager

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

var (
	podCIDRTypeField     *field.Path = field.NewPath("spec").Child("podCIDRType")
	extraCIDRField       *field.Path = field.NewPath("spec").Child("extraCIDR")
	podMACPrefixField    *field.Path = field.NewPath("spec").Child("podMACPrefix")
	podRPFilterField     *field.Path = field.NewPath("spec").Child("podRPFilter")
	txQueueLenField      *field.Path = field.NewPath("spec").Child("txQueueLen")
	vethLinkAddressField *field.Path = field.NewPath("spec").Child("vethLinkAddress")
)

func validateCreateCoordinator(coord *spiderpoolv2beta1.SpiderCoordinator) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateUpdateCoordinator(oldCoord, newCoord *spiderpoolv2beta1.SpiderCoordinator) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func ValidateCoordinatorSpec(spec *spiderpoolv2beta1.CoordinatorSpec, requireOptionalType bool) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateCoordinatorPodCIDRType(t string) *field.Error { _ = "STUB: not implemented"; return nil }

func validateCoordinatorExtraCIDR(cidrs []string) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateCoordinatorPodMACPrefix(prefix *string) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateCoordinatorPodRPFilter(f *int) *field.Error { _ = "STUB: not implemented"; return nil }
