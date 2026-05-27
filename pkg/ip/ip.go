// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ip

import (
	"math/big"
	"net"
	"net/netip"

	"github.com/spidernet-io/spiderpool/pkg/types"
)

// IsIPVersion reports whether version is a valid IP version (4 or 6).
func IsIPVersion(version types.IPVersion) error { _ = "STUB: not implemented"; return nil }

// ParseIP parses IP string as a CIDR notation IP address of the specified
// IP version.
func ParseIP(version types.IPVersion, s string, isCIDR bool) (*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ContainsIP reports whether the subnet parsed from the subnet string
// includes the IP address parsed from the IP string. Both must belong
// to the same IP version.
func ContainsIP(version types.IPVersion, subnet string, ip string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsIP reports whether IP string is a IP address of the specified IP version.
func IsIP(version types.IPVersion, s string) error { _ = "STUB: not implemented"; return nil }

// IPsDiffSet calculates the difference set of two IP address slices.
// For example, the difference set between [172.18.40.1 172.18.40.2] and
// [172.18.40.2 172.18.40.3] is [172.18.40.1].
//
// If sorted is true, the result set of IP addresses will be sorted.
func IPsDiffSet(ipSourceList, ipExcludeList []net.IP, sorted bool) []net.IP {
	_ = "STUB: not implemented"
	return nil
}

func IsDiffIPSet(ipSourceList, ipExcludeList []net.IP) bool {
	_ = "STUB: not implemented"
	return false
}

// getIPDiffSet returns a list of IPs from ipSourceList that are not in ipExcludeList. Parameters:
// - ipSourceList: a slice of net.IP that represents the source list of IPs.
// - ipExcludeList: a slice of net.IP that represents the list of IPs to be excluded.
// - sorted: a boolean indicating whether the resulting list should be sorted.
// - expectCount: an integer specifying the maximum number of IPs to return. If expectCount <= 0, all IPs will be returned.
func getIPDiffSet(ipSourceList, ipExcludeList []net.IP, sorted bool, expectCount int) []net.IP {
	_ = "STUB: not implemented"
	return nil
}

// FindAvailableIPs find available ip list in range
func FindAvailableIPs(ipRanges []string, ipList []net.IP, count int) []net.IP {
	_ = "STUB: not implemented"
	return nil
}

func nextIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// IPsUnionSet calculates the union set of two IP address slices.
// For example, the union set between [172.18.40.1 172.18.40.2] and
// [172.18.40.2 172.18.40.3] is [172.18.40.1 172.18.40.2 172.18.40.3].
//
// If sorted is true, the result set of IP addresses will be sorted.
func IPsUnionSet(ips1, ips2 []net.IP, sorted bool) []net.IP { _ = "STUB: not implemented"; return nil }

// IPsIntersectionSet calculates the intersection set of two IP address
// slices. For example, the intersection set between [172.18.40.1 172.18.40.2]
// and [172.18.40.2 172.18.40.3] is [172.18.40.2].
//
// If sorted is true, the result set of IP addresses will be sorted.
func IPsIntersectionSet(ips1, ips2 []net.IP, sorted bool) []net.IP {
	_ = "STUB: not implemented"
	return nil
}

// NextIP returns the next IP address.
func NextIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// PrevIP returns the previous IP address.
func PrevIP(ip net.IP) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// Cmp compares two IP addresses, returns according to the following rules:
// ip1 < ip2: -1
// ip1 = ip2: 0
// ip1 > ip2: 1
func Cmp(ip1, ip2 net.IP) int { _ = "STUB: not implemented"; return 0 }

// ipToInt converts net.IP to big.Int.
func ipToInt(ip net.IP) *big.Int { _ = "STUB: not implemented"; return nil }

// intToIP converts big.Int to net.IP.
func intToIP(i *big.Int) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func ParseIPOrCIDR(s string) (netip.Prefix, error) {
	_ = "STUB: not implemented"
	return *new(netip.Prefix), nil
}
