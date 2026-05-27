// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"net"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
)

type coordinator struct {
	firstInvoke                                                  bool
	ipFamily, currentRuleTable, hostRuleTable                    int
	tuneMode                                                     Mode
	hostVethName, podVethName, vethLinkAddress, currentInterface string
	v4HijackRouteGw, v6HijackRouteGw                             net.IP
	HijackCIDR                                                   []string
	netns, hostNs                                                ns.NetNS
	hostVethHwAddress, podVethHwAddress                          net.HardwareAddr
	currentAddress                                               []netlink.Addr
	v4PodOverlayNicAddr, v6PodOverlayNicAddr                     *net.IPNet
	hostIPRouteForPod                                            []net.IP
}

func (c *coordinator) autoModeToSpecificMode(mode Mode, podFirstInterface string, vethExist bool) error {
	_ = "STUB: not implemented"
	return nil
}

// firstInvoke check if coordinator is first called and do some checks:
// underlay mode only works with underlay mode, which can't work with overlay
// mode, and which can't be called in first cni invoked by using multus's
// annotations: v1.multus-cni.io/default-network
func (c *coordinator) coordinatorModeAndFirstInvoke(logger *zap.Logger, podFirstInterface string) error {
	_ = "STUB: not implemented"
	return nil
}

// underlay mode can't work with calico/cilium(overlay)

// ensure that each NIC has a separate policy routing table number

// keep table 100 for eth0, first non-eth0 nic is table 101

// for non-eth0 or non first-underlay nic, Policy routing
// table numbers are cumulative based on the number of NICs
// for example:
// there are veth0, eth0,net1,net2 nic, the policy routing table numbers
// of net2 is:  4 + 98 == 102.

// in overlay mode, it should no veth0 and currentInterface isn't eth0

// if pod has only eth0 and net1, the first invoke is true

// keep table 100 for eth0, first non-eth0 nic is table 101

// for non-eth0 or non first-underlay nic, Policy routing
// table numbers are cumulative based on the number of NICs
// for example:
// there are eth0,net1,net2 nic, the policy routing table numbers
// of net2 is:  3 + 99 == 102.

func (c *coordinator) checkNICState(iface string) error { _ = "STUB: not implemented"; return nil }

// setupVeth sets up a pair of virtual ethernet devices. move one to the host and other
// one to container.
func (c *coordinator) setupVeth(logger *zap.Logger, containerID string) error {
	_ = "STUB: not implemented"
	// systemd 242+ tries to set a "persistent" MAC addr for any virtual device
	// by default (controlled by MACAddressPolicy). As setting happens
	// asynchronously after a device has been created, ep.Mac and ep.HostMac
	// can become stale which has a serious consequence - the kernel will drop
	// any packet sent to/from the endpoint. However, we can trick systemd by
	// explicitly setting MAC addrs for both veth ends. This sets
	// addr_assign_type for NET_ADDR_SET which prevents systemd from changing
	// the addrs.
	return nil
}

// set an address to veth to fix work with istio
// set only when not ipv6 only

// setupNeighborhood setup neighborhood tables for pod and host.
// equivalent to: `ip neigh add ....`
func (c *coordinator) setupNeighborhood(logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// do any cleans?

// setupRoutes setup hijack subnet routes for pod and host
// equivalent to: `ip route add $route table $ruleTable`
func (c *coordinator) setupHijackRoutes(logger *zap.Logger, ruleTable int) error {
	_ = "STUB: not implemented"
	return nil
}

// make sure that veth0/eth0 forwards traffic within the cluster
// eq: ip route add <cluster/service cidr> dev veth0/eth0

// setupHostRoutes create routes for all host IPs, make sure that traffic to
// pod's host is forward to veth pair device.
func (c *coordinator) setupHostRoutes(logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// traffic sent to the pod its node is forwarded via veth0/eth0
// eq: "ip r add <ipAddressOnNode> dev veth0/eth0 table <ruleTable>"

// make sure `ip rule from all lookup 500 pref 32765` exist

// do any cleans dirty route tables

// set routes for host
// equivalent: ip add  <chainedIPs> dev <hostVethName> table  on host

// tunePodRoutes make sure that move all routes of podDefaultRouteNIC interface to main table, and move original routes
// in main table to new table
func (c *coordinator) tunePodRoutes(logger *zap.Logger, configDefaultRouteNIC string) error {
	_ = "STUB: not implemented"
	return nil
}

// the current interface's default route no found, we can keep all routes of
// this nic in main table, and don't tune the routes

// configDefaultRouteNIC is empty by default, and we always keep the all routes of the
// first NIC is in main and move the all routes of non-first NIC to policy routing table.
// see https://github.com/spidernet-io/spiderpool/issues/2176.

// make sure that traffic sent from current interface to lookup table <ruleTable>
// eq: ip rule add from <currentInterfaceIPAddress> lookup <ruleTable>

// mv calico or cilium default route to table 100 to fix to the problem of
// inconsistent routes, the pod forwards the response packet from net1 (macvlan)
// when it sends the response packet. but the request packet comes from eth0(calico).
// see https://github.com/spidernet-io/spiderpool/issues/3683

// copy to table 100

// move all routes of the specified interface to a new route table

// move all routes of the specified interface to a new route table

// makeReplyPacketViaVeth make sure that tcp replay packet is forward by veth0
// NOTE: underlay mode only.
func (c *coordinator) makeReplyPacketViaVeth(logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// getHostVethName select the first 11 characters of the containerID for the host veth.
func getHostVethName(containerID string) string { _ = "STUB: not implemented"; return "" }

func min(len int) int { _ = "STUB: not implemented"; return 0 }

func getMarkInt(markBit int) int { _ = "STUB: not implemented"; return 0 }

func getMarkString(mark int) string { _ = "STUB: not implemented"; return "" }

func (c *coordinator) ensureIPtablesRule(iptablesInterfaces []utiliptables.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

func GetAllHostIPRouteForPod(c *coordinator, ipFamily int, allPodIP []netlink.Addr) (finalNodeIPList []net.IP, e error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get node ip by `ip r get podIP`

// get additional host ip

func (c *coordinator) AnnounceIPs(logger *zap.Logger) error { _ = "STUB: not implemented"; return nil }

// send an gratuitous arp to announce the new mac address
