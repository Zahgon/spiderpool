// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ip

import (
	"github.com/spidernet-io/spiderpool/pkg/types"
)

// IsRoute reports whether dst and gw strings constitute a route of the
// specified IP version.
func IsRoute(version types.IPVersion, dst, gw string) error { _ = "STUB: not implemented"; return nil }

// IsRouteWithoutIPVersion reports whether dst and gw strings constitute
// an route.
func IsRouteWithoutIPVersion(dst, gw string) error { _ = "STUB: not implemented"; return nil }

// IsIPv4Route reports whether dst and gw strings constitute an IPv4 route.
func IsIPv4Route(dst, gw string) bool { _ = "STUB: not implemented"; return false }

// IsIPv6Route reports whether dst and gw strings constitute an IPv6 route.
func IsIPv6Route(dst, gw string) bool { _ = "STUB: not implemented"; return false }
