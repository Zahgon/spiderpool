// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ip

import (
	"net"

	"github.com/spidernet-io/spiderpool/pkg/types"
)

// ParseCIDR parses subnet string as a CIDR notation IP address of the
// specified IP version, like "172.18.40.0/24" or "fd00:db8::/32".
func ParseCIDR(version types.IPVersion, subnet string) (*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ContainsCIDR reports whether subnet1 includes subnet2. Both of them
// are parsed from subnet strings and must belong to the same IP version.
func ContainsCIDR(version types.IPVersion, subnet1 string, subnet2 string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsCIDROverlap reports whether the subnets of specific IP version
// parsed from two subnet strings overlap.
func IsCIDROverlap(version types.IPVersion, subnet1, subnet2 string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func containsCIDR(subnet1 string, subnet2 string) bool {
	_ = "STUB: not implemented"
	// Ignore the error returned here. The format of the subnet should be
	// verified in external IsCIDR.
	return false
}

// IsCIDR reports whether subnet string is a CIDR notation IP address
// of the specified IP version.
func IsCIDR(version types.IPVersion, subnet string) error { _ = "STUB: not implemented"; return nil }

// IsIPv4CIDR reports whether subnet string is a CIDR notation IPv4 address.
func IsIPv4CIDR(subnet string) bool { _ = "STUB: not implemented"; return false }

// IsIPv6CIDR reports whether subnet string is a CIDR notation IPv6 address.
func IsIPv6CIDR(subnet string) bool { _ = "STUB: not implemented"; return false }

func IsFormatCIDR(subnet string) error { _ = "STUB: not implemented"; return nil }
