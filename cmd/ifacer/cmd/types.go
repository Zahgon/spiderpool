// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/containernetworking/cni/pkg/types"
	"github.com/vishvananda/netlink"
)

var DefaultBondName = "sp_bond0"

type Ifacer struct {
	types.NetConf
	Interfaces []string `json:"interfaces,omitempty"`
	VlanID     int      `json:"vlanID,omitempty"`
	Bond       *Bond    `json:"bond,omitempty"`
}

type Bond struct {
	Name    string `json:"name,omitempty"`
	Mode    int    `json:"mode,omitempty"`
	Options string `json:"options,omitempty"`
}

func ParseConfig(stdin []byte) (*Ifacer, error) { _ = "STUB: not implemented"; return nil, nil }

func validateBondMode(mode int) error { _ = "STUB: not implemented"; return nil }

// parseBondOptions convert options string to BondOptions object.
// incorrect options are ignored without an error return
// options input-formatted: "k1=v1;k2=v2;k3=v3"
func parseString2BondOptions(options string) (*BondOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseBondOptions2NetlinkBond convert BondOptions to netlink.Bond object
func parseBondOptions2NetlinkBond(bondOptions *BondOptions, bond *netlink.Bond) error {
	_ = "STUB: not implemented"
	return nil
}
