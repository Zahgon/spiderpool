// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package sysctl

import (
	"github.com/containernetworking/plugins/pkg/ns"
)

var (
	SysctlRPFilter   = "net.ipv4.conf.all.rp_filter"
	SysctlEnableIPv6 = "net.ipv6.conf.all.disable_ipv6"
)

// DefaultSysctlConfig is the default sysctl config for the node
var DefaultSysctlConfig = []struct {
	Name           string
	Value          string
	IsIPv4, IsIPv6 bool
}{
	// In order to avoid large-scale cluster arp_table overflow, resulting in
	// pods not being able to communicate or pods not being able to start due
	// to the inability to insert static arp table entries, it is necessary
	// to appropriately increase and adjust its value. more details see:
	// https://github.com/spidernet-io/spiderpool/issues/3587
	{
		Name: "net.ipv4.neigh.default.gc_thresh3",
		// Assuming a node is full of underlay pods (110) and their subnet
		// mask is 16 bits ( 2 ^ 8 = 256 IPs), the value is 110 * 256 = 28160
		Value:  "28160",
		IsIPv4: true,
	},
	{
		// this sysctl may not be available at low kernel levels,
		// so we'll ignore it at this point.
		Name:   "net.ipv6.neigh.default.gc_thresh3",
		Value:  "28160",
		IsIPv6: true,
	},
	// send gratitous ARP when device or address change
	{
		Name:   "net.ipv4.conf.all.arp_notify",
		Value:  "1",
		IsIPv4: true,
	},
	{
		Name:   "net.ipv4.conf.all.forwarding",
		Value:  "1",
		IsIPv4: true,
	},
	{
		Name:   "net.ipv6.conf.all.forwarding",
		Value:  "1",
		IsIPv6: true,
	},
	{
		Name:   "net.ipv4.conf.all.rp_filter",
		Value:  "0",
		IsIPv4: true,
		IsIPv6: true,
	},
}

// SetSysctlRPFilter sets rp_filter value for host netns and specify netns.
func SetSysctlRPFilter(netns ns.NetNS, value int32) error {
	_ = "STUB: not implemented"
	// set pod rp_filter
	return nil
}

// EnableIpv6Sysctl enable ipv6 for specify netns
func EnableIpv6Sysctl(value int32) error { _ = "STUB: not implemented"; return nil }

func SetSysctl(sysConfig string, value string) error {
	_ = "STUB: not implemented"
	// sysConfig: net.ipv6.neigh.default.gc_thresh3
	// to: net/ipv6/neigh/default/gc_thresh3
	return nil
}

func EnableIPv6ForInterfaces(ifaces []string) error { _ = "STUB: not implemented"; return nil }

// Read current sysctl value

// Write sysctl to enable IPv6
