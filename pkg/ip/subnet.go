// Copyright 2024 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ip

import (
	"math/big"
	"net"
)

type Range struct {
	Raw        string
	Start, End *big.Int
}

type CIDR struct {
	base    net.IPNet           // base is the base subnet
	isIPv4  bool                // isIPv4 is true if the base subnet is an IPv4 address
	ranges  []Range             // ranges is available IP ranges
	usedMap map[string]struct{} // usedMap is a map of used IP addresses
}

func NewCIDR(base string, ranges []string, exclude []string) (*CIDR, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddUsedIP adds the IP addresses to the used store.
func (c *CIDR) AddUsedIP(list ...string) error { _ = "STUB: not implemented"; return nil }

// TotalUsedIPInt returns the number of used IP addresses in the subnet.
func (c *CIDR) TotalUsedIPInt() int { _ = "STUB: not implemented"; return 0 }

// TotalIP returns the number of IP addresses in the subnet.
func (c *CIDR) TotalIP() *big.Int { _ = "STUB: not implemented"; return nil }

// TotalIPInt returns the number of IP addresses in the subnet as an integer.
func (c *CIDR) TotalIPInt() int { _ = "STUB: not implemented"; return 0 }

// addIncludeRange adds the range to the subnet.
func (c *CIDR) addIncludeRange(r Range) error { _ = "STUB: not implemented"; return nil }

// if start > end, start = end, end = start

// a.start in b

// a.end in b

// b.start in a

// b.end in a

// parse parses the range string and returns the start and end IP addresses.
// case ipv4
// - "10.6.0.1"
// - "10.6.0.1-10.6.0.1"
// - "10.6.0.1/24"
// case ipv6
// - "fd00:db8::1"
// - "fd00:db8::1-fd00:db8::1"
// - "fd00:db8::1/64"
func (c *CIDR) parse(value string) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// convertIPtoBigInt converts the IP address to big.Int.
func (c *CIDR) convertIPtoBigInt(value string) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IPRange returns the IP ranges in the subnet.
func (c *CIDR) IPRange() []Range {
	_ = "STUB: not implemented"

	// IsOverlapIPRanges checks if the subnet overlaps with the given ranges.
	return nil
}

func (c *CIDR) IsOverlapIPRanges(r []Range) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func removeExcludeIPRange(base []Range, exclude []Range) []Range {
	_ = "STUB: not implemented"
	return nil
}
