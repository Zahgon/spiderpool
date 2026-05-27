// Copyright 2025 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package networking

import (
	"net"
	"time"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/mdlayher/ndp"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
)

var (
	retryNum = 3
	timeOut  = 100 * time.Millisecond
)

type Detector struct {
	logger                                                                   *zap.Logger
	enableIPv4ConflictDetection, enableIPv6ConflictDetection                 bool
	enableIPv4GatewayReachableDetection, enableIPv6GatewayReachableDetection bool
	retries                                                                  int
	iface                                                                    string
	timeout                                                                  time.Duration
	ip4, ip6, v4Gw, v6Gw                                                     net.IP
}

func DetectIPConflictAndGatewayReachable(logger *zap.Logger, iface string, hostNs ns.NetNS, netns ns.NetNS, ipconfigs []*models.IPConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// spiderpool assigns IPs to all NICs in advance of the first call to ipam.
// different NICs come from different pools, so we only need to focus on the current NIC's ipconfig.

// IP conflict detection and gateway detection are disabled

func (d *Detector) ARPDetect() error { _ = "STUB: not implemented"; return nil }

// IP conflict detection must precede gateway detection, which avoids the
// possibility that gateway detection may update arp table entries first and cause
// communication problems when IP conflict detection fails
// see https://github.com/spidernet-io/spiderpool/issues/4475
// call ip conflict detection

//  we do detect gateway connection lastly
// Finally, there is gateway detection, which updates the correct arp table entries
// once there are no IP address conflicts and fixed Mac addresses
// call gateway detection

func (d *Detector) NDPDetect() error { _ = "STUB: not implemented"; return nil }

// wait for ndp ready

// When DAD(Duplicate Address Detection) is enanled, the kernel will check if this local link address is in conflict,
// this may take a while, set the maximum timeout to 10s

// IP conflict detection must precede gateway detection, which avoids the
// possibility that gateway detection may update arp table entries first and cause
// communication problems when IP conflict detection fails
// see https://github.com/spidernet-io/spiderpool/issues/4475
// call ip conflict detection

// we do detect gateway connection lastly
// Finally, there is gateway detection, which updates the correct arp table entries
// once there are no IP address conflicts and fixed Mac addresses
// call gateway detection

func (d *Detector) detectIP4Conflicting(l netlink.Link, ifi *net.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

// Set a timeout of d.timeout for receiving packets

// we send a gratuitous arp to checking if ip is conflict
// we use dad mode(duplicate address detection mode), so
// we set source ip to 0.0.0.0

// For some edge cases, even if we set the ReadTimeOut for the ARP connection,
// it may not take effect. The arpClient.Read function keeps receiving unexpected errors,
// causing the entire for loop to be unable to exit.

// Read a packet from the socket.

// found ip conflicting

// If an arp reply is not received within the timeout period or is not
// a expected arp reply

// unexpected error, retrying send and receive...

func (d *Detector) detectGateway4Reachable(l netlink.Link, ifi *net.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

// Set a timeout of d.timeout for receiving packets

// Read a packet from the socket.

// For some edge cases, even if we set the ReadTimeOut for the ARP connection,
// it may not take effect. The arpClient.Read function keeps receiving unexpected errors,
// causing the entire for loop to be unable to exit.

// Now we catch an ARP response

// Check if the sender's MAC address is the same as the interface's address

// If an arp reply is not received within the timeout period or is not
// sent from the gateway IP, it is assumed that the gateway is not reachable.

func (d *Detector) detectIP6Conflicting(ifi *net.Interface, ndpClient *ndp.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

// no ndp response unitil timeout, indicates gateway unreachable

// retry it if is other error

func (d *Detector) detectGateway6Reachable(ifi *net.Interface, ndpClient *ndp.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

// no ndp response unitil timeout, indicates gateway unreachable

// retry it if is other error
