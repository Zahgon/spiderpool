// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package common

const (
	ENV_INSTALL_OVERLAY          = "INSTALL_OVERLAY_CNI"
	E2E_SPIDERPOOL_ENABLE_SUBNET = "E2E_SPIDERPOOL_ENABLE_SUBNET"
	INSTALL_CALICO               = "INSTALL_CALICO"
	INSTALL_CILIUM               = "INSTALL_CILIUM"
	ENABLE_DRA                   = "E2E_SPIDERPOOL_ENABLE_DRA"
)

func checkBoolEnv(name string) bool { _ = "STUB: not implemented"; return false }

func CheckRunOverlayCNI() bool { _ = "STUB: not implemented"; return false }

func IsDRAEnabled() bool { _ = "STUB: not implemented"; return false }

func CheckSubnetFeatureOn() bool { _ = "STUB: not implemented"; return false }

func CheckCalicoFeatureOn() bool { _ = "STUB: not implemented"; return false }

func CheckCiliumFeatureOn() bool { _ = "STUB: not implemented"; return false }
