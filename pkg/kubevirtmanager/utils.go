// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package kubevirtmanager

import (
	kubevirtv1 "kubevirt.io/api/core/v1"
)

func isVMIControlledByVM(vmi *kubevirtv1.VirtualMachineInstance) bool {
	_ = "STUB: not implemented"
	return false
}
