// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package networking

import (
	"net"
	"net/netip"

	"github.com/containernetworking/plugins/pkg/ns"
	"go.uber.org/zap"
)

// OverwriteHwAddress override the hardware address of the specified interface.
func OverwriteHwAddress(logger *zap.Logger, netns ns.NetNS, macPrefix, iface string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// we only focus on first element

// newmac = xx:xx + xx:xx:xx:xx

// parseMac parse hardware addr from given string
func parseMac(s string) net.HardwareAddr { _ = "STUB: not implemented"; return *new(net.HardwareAddr) }

// inetAton converts an IP Address (IPv4 or IPv6) netip.addr object to a hexadecimal representation.
// for ipv4: convert a full IP address(length: 4 B) to hexadecimal representation.
// for ipv6: convert
func inetAton(ip netip.Addr) (string, error) { _ = "STUB: not implemented"; return "", nil }

// 32 bit -> 4 B

// for ipv6: 128 bit = 32 hex
// take the last 8 hex as the hardware address

// convertHex2Mac convert hexcode to 4B hardware address
// convert ip(hex) to "xx:xx:xx:xx"
func convertHex2Mac(hexCode []byte) string { _ = "STUB: not implemented"; return "" }

// GetHwAddressByName get hardware address of veth pair device
func GetHwAddressByName(netns ns.NetNS, podVethPairName, hostVethPairName string) (net.HardwareAddr, net.HardwareAddr, error) {
	_ = "STUB: not implemented"
	return *new(net.HardwareAddr), *new(net.HardwareAddr), nil
}

// GetHostVethName get veth name in host side
func GetHostVethName(netns ns.NetNS, podVethPairName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// get link index of host veth-peer and pod veth-peer mac-address

// AddStaticNeighborTable add a static neighborhood table
func AddStaticNeighborTable(linkIndex int, dstIP net.IP, hwAddress net.HardwareAddr) error {
	_ = "STUB: not implemented"
	return nil
}
