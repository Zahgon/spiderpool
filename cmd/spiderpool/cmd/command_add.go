// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/containernetworking/cni/pkg/skel"
	current "github.com/containernetworking/cni/pkg/types/100"

	"github.com/spidernet-io/spiderpool/api/v1/agent/client/daemonset"
)

// CmdAdd follows CNI SPEC cmdAdd.
func CmdAdd(args *skel.CmdArgs) (err error) { _ = "STUB: not implemented"; return nil }

// Defer a panic recover, so that in case we panic we can still return
// a proper error to the runtime.

// If it is recovering and an error occurs, then we need to
// present both.

// When IPAM is invoked, the NIC is down and must be set it up in order to detect IP conflicts and
// gateway reachability.

// Validate IPAM request response.

// do ip conflict and gateway detection

// CNI will set the interface to up, and the kernel only sends GARPs/Unsolicited NA when the interface
// goes from down to up or when the link-layer address changes on the interfaces. in order to the
// kernel send GARPs/Unsolicited NA when the interface goes from down to up.
// see https://github.com/spidernet-io/spiderpool/issues/4650

// Assemble the result of IPAM request response.

// assembleResult groups the IP allocation resutls of IPAM request response
// based on NIC and combines them into CNI results.
func assembleResult(cniVersion, IfName string, ipamResponse *daemonset.PostIpamIPOK) (*current.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mock DNS.
