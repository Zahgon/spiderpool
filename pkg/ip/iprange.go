// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ip

import (
	"net"

	"github.com/spidernet-io/spiderpool/pkg/types"
)

// MergeIPRanges merges dispersed IP ranges.
// For example, transport [172.18.40.1-172.18.40.3, 172.18.40.2-172.18.40.5]
// to [172.18.40.1-172.18.40.5]. The overlapping part of two IP ranges will
// be ignored.
func MergeIPRanges(version types.IPVersion, ipRanges []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseIPRanges parses IP ranges as a IP address slices of the specified
// IP version.
func ParseIPRanges(version types.IPVersion, ipRanges []string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseIPRange parses IP range as an IP address slices of the specified
// IP version.
func ParseIPRange(version types.IPVersion, ipRange string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 'n' must be 1 or 2 because of the validation of 'IsIPRange'

// ConvertIPsToIPRanges converts the IP address slices of the specified
// IP version into a group of distinct, sorted and merged IP ranges.
func ConvertIPsToIPRanges(version types.IPVersion, ips []net.IP) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ContainsIPRange reports whether the subnet parsed from the subnet string
// includes the IP address slices parsed from the IP range. Both must belong
// to the same IP version.
func ContainsIPRange(version types.IPVersion, subnet string, ipRange string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IPRangeContainsIP reports whether the IP range includes the IP address.
// Both must belongto the same IP version.
func IPRangeContainsIP(version types.IPVersion, ipRange string, ip string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsIPRangeOverlap reports whether the IP address slices of specific IP
// version parsed from two IP ranges overlap.
func IsIPRangeOverlap(version types.IPVersion, ipRange1, ipRange2 string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Ignore the error returned here. The format of the IP range has been
// verified in IsIPRange above.

// IsIPRange reports whether ipRange string is a valid IP range. An IP
// range can be a single IP address in the style of '172.18.40.0', or
// an address range in the form of '172.18.40.0-172.18.40.10'.
// The following formats are invalid:
// "172.18.40.0 - 172.18.40.10": there can be no space between two IP
// addresses.
// "172.18.40.1-2001:db8:a0b:12f0::1": invalid combination of IPv4 and
// IPv6.
// "172.18.40.10-172.18.40.1": the IP range must be ordered.
func IsIPRange(version types.IPVersion, ipRange string) error {
	_ = "STUB: not implemented"
	return nil
}

// IsIPv4IPRange reports whether ipRange string is a valid IPv4 range.
// See IsIPRange for more description of IP range.
func IsIPv4IPRange(ipRange string) bool { _ = "STUB: not implemented"; return false }

// IsIPv6IPRange reports whether ipRange string is a valid IPv6 range.
// See IsIPRange for more description of IP range.
func IsIPv6IPRange(ipRange string) bool { _ = "STUB: not implemented"; return false }
