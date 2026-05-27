// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/vishvananda/netlink"
)

func CmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// create vlan interface

func createBondDevice(conf *Ifacer) (*netlink.Bond, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create vlan interface base on bond

func createVlanDevice(conf *Ifacer) error {
	_ = "STUB: not implemented"

	// If the parent interface is down, we set it to up.
	return nil
}

// we only create if vlanIf not present
