// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package networking

import (
	"net"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"
	"go.uber.org/zap"
)

var defaultRulePriority = 1000

// GetRoutesByName return all routes is belonged to specify interface
// filter by family also
func GetRoutesByName(iface string, ipfamily int) (routes []netlink.Route, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetDefaultGatewayByName(iface string, ipfamily int) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddToRuleTable(dst *net.IPNet, ruleTable int) error { _ = "STUB: not implemented"; return nil }

func DelToRuleTable(dst *net.IPNet, ruleTable int) error { _ = "STUB: not implemented"; return nil }

func AddRuleTableWithMark(mark, ruleTable, ipFamily int) error {
	_ = "STUB: not implemented"
	return nil
}

// AddFromRuleTable add route rule for calico/cilium cidr(ipv4 and ipv6)
// Equivalent to: `ip rule add from <cidr> `
func AddFromRuleTable(src *net.IPNet, ruleTable int) error { _ = "STUB: not implemented"; return nil }

// DelFromRuleTable equivalent to: `ip rule del from <cidr> lookup <ruletable>`
func DelFromRuleTable(src *net.IPNet, ruleTable int) error { _ = "STUB: not implemented"; return nil }

// AddRoute add static route to specify rule table
func AddRoute(logger *zap.Logger, ruleTable, ipFamily int, scope netlink.Scope, iface string, src, dst *net.IPNet, v4Gw, v6Gw net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

func GetLinkIndexAndRoutes(iface string, ipfamily int) (int, []netlink.Route, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// CopyDefaultRoute found the default route of pod's eth0 nic, and copy this
// to dstRuleTable.
func CopyDefaultRoute(logger *zap.Logger, iface string, srcRuleTable, podOverlayDefaultRouteRuleTable, ipfamily int) error {
	_ = "STUB: not implemented"
	return nil
}

// only handle route tables from table main

// ignore local link route

// MoveRouteTable move all routes of the specified interface to a new route table
// Equivalent: `ip route del <route>` and `ip r route add <route> <table>`
func MoveRouteTable(logger *zap.Logger, iface string, srcRuleTable, dstRuleTable, ipfamily int) error {
	_ = "STUB: not implemented"
	return nil
}

// only handle route tables from table main

// ignore local link route

// moveRouteTable move route table from srcRuleTable to dstRuleTable. NOTE: if copyOverlayDefaultRoute is true,
// only add the default route to host rule table and exit in advance.
func moveRouteTable(linkIndex, srcRuleTable, dstRuleTable int, onlyCopyOverlayDefaultRoute bool, route netlink.Route, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// only copy overlay default route, don't need delete the default route

// Del the default route from main

// only copy overlay default route, don't need add non-default routes

// we need copy the all routes in main table of the podDefaultRouteNic to dstRuleTable.
// Otherwise, we don't know how to forward the packet send from the nic

// in high kernel, if pod has multi ipv6 default routes, all default routes
// will be put in MultiPath
/*
	{
		Gw: [{Ifindex: 3 Weight: 1 Gw: fd00:10:7::103 Flags: []} {Ifindex: 5 Weight: 1 Gw: fd00:10:6::100 Flags: []}]}"
	}
*/

// get generated default Route for new table

// GetDefaultRouteInterface returns the name of the NIC where the default route is located
// if filterInterface not be empty, return first default route interface
// otherwise filter filterInterface
func GetDefaultRouteInterface(ipfamily int, filterInterface string, netns ns.NetNS) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ConvertMaxMaskIPNet(nip net.IP) *net.IPNet { _ = "STUB: not implemented"; return nil }
