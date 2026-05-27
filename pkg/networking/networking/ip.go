// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package networking

import (
	"net"

	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
)

// GetIPFamilyByResult return IPFamily by parse CNI Result
func GetIPFamilyByResult(prevResult *current.Result) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetGatewayIP(addrs []netlink.Addr) (v4Gw, v6Gw net.IP, err error) {
	_ = "STUB: not implemented"
	return *new(net.IP), *new(net.IP), nil
}

// IPAddressByName returns all IP addresses of the given pod's interface
// group by ipFamily
func IPAddressByName(netns ns.NetNS, interfacenName string, ipFamily int) ([]netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IPAddressOnNode return all ip addresses on the node, filter by ipFamily
// skipping any interfaces whose name matches any of the exclusion list regexes
func GetAllIPAddress(ipFamily int, excludeInterface []string) ([]netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAddersByName return all unicast ip address of interface, filter by ipFamily
func GetAddersByName(iface string, ipfamily int) ([]netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAddersByLink return all unicast ip address of interface, filter by interface,ipFamily
func GetAddersByLink(link netlink.Link, ipfamily int) ([]netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getAdders(link netlink.Link, ipfamily int) ([]netlink.Addr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CheckInterfaceExist(netns ns.NetNS, iface string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func isInterfaceExist(iface string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func GetUPLinkList(netns ns.NetNS) ([]netlink.Link, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only include devices in up|broadcast|multicast state
// include:
//   example: eth0/net1, flag: up|broadcast|multicast
// exclude:
//   lo, flags: up|loopback
//   tunl0: flags: 0

func LinkSetBondSlave(slave string, bond *netlink.Bond) error {
	_ = "STUB: not implemented"
	return nil
}

func LinkSetTxqueueLen(iface string, txQueneLen int) error { _ = "STUB: not implemented"; return nil }

func LinkAdd(link netlink.Link) error { _ = "STUB: not implemented"; return nil }

func linkAddAndSetUp(link netlink.Link) error { _ = "STUB: not implemented"; return nil }

// IPNetEqual returns true iff both IPNet are equal
// Copyright Authors of vishvananda/netlink
func IPNetEqual(ipn1 *net.IPNet, ipn2 *net.IPNet) bool { _ = "STUB: not implemented"; return false }
