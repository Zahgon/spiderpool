// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ip

import (
	"net"

	"github.com/spidernet-io/spiderpool/pkg/types"
)

func AssembleTotalIPs(ipVersion types.IPVersion, ipRanges, excludedIPRanges []string) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CIDRToLabelValue(ipVersion types.IPVersion, subnet string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
