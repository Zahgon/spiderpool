// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/containernetworking/cni/pkg/skel"
)

func CmdAdd(args *skel.CmdArgs) (err error) { _ = "STUB: not implemented"; return nil }

// parse prevResult

// checking if the nic is in up state

// check if it's first time invoke

// get all ip of pod

// get ip addresses of the node

// get basic info

// get ips of this interface(preInterfaceName) from, including ipv4 and ipv6

// Fixed Mac addresses must come after IP conflict detection, otherwise the switch learns to communicate
// with the wrong Mac address when IP conflict detection fails

// set txqueuelen

// =================================

// ensure ipv6 is enable

// for ipv6, maybe kernel require the src must be from the device in
// some kernel version.
// refer to https://bugzilla.kernel.org/show_bug.cgi?id=107071

// get v4 and v6 gw for hijick route'gw
