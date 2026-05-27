// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"net"

	"github.com/vishvananda/netlink"
)

// BondOptions  for the bonding driver are supplied as parameters to the
// bonding module at load time, or are specified via sysfs.
// refer to https://www.kernel.org/doc/Documentation/networking/bonding.txt
type BondOptions struct {
	ActiveSlave     string   `json:"active_slave,omitempty"`
	AdActorSysPrio  int      `json:"ad_actor_sys_prio,omitempty"`
	AdActorSystem   string   `json:"ad_actor_system,omitempty"`
	AdSelect        int      `json:"ad_select,omitempty"`
	AdUserPortKey   int      `json:"ad_user_port_key,omitempty"`
	AllSlavesActive int      `json:"all_slaves_active,omitempty"`
	ArpInterval     int      `json:"arp_interval,omitempty"`
	ArpIPTargets    []string `json:"arp_ip_target,omitempty"`
	ArpValidate     int      `json:"arp_validate,omitempty"`
	ArpAllTargets   int      `json:"arp_all_targets,omitempty"`
	DownDelay       int      `json:"downdelay,omitempty"`
	FailOverMac     int      `json:"fail_over_mac,omitempty"`
	LacpRate        int      `json:"lacp_rate,omitempty"`
	Miimon          int      `json:"miimon,omitempty"`
	MinLinks        int      `json:"min_links,omitempty"`
	PacketsPerSlave int      `json:"packets_per_slave,omitempty"`
	// The primary option is only valid for active-backup(1),
	// balance-tlb (5) and balance-alb (6) mode
	Primary         string `json:"primary,omitempty"`
	PrimaryReselect int    `json:"primary_reselect,omitempty"`
	TlbDynamicLb    int    `json:"tlb_dynamic_lb,omitempty"`
	UpDelay         int    `json:"up_delay,omitempty"`
	UseCarrier      int    `json:"use_carrier,omitempty"`
	XmitHashPolicy  int    `json:"xmit_hash_policy,omitempty"`
	LpInterval      int    `json:"lp_interval,omitempty"`
	ResendIgmp      int    `json:"resend_igmp,omitempty"`
	NumPeerNotif    int    `json:"peer_notif_delay,omitempty"`
}

type BondOptionFunc func(bond *netlink.Bond)

func ActiveSlaveOption(activeSlave int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func AdActorSystemOption(adActorSystem net.HardwareAddr) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func AdActorSysPrioOption(adActorSysPrio int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func AdSelectOption(adSelect int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func AdUserPortKeyOption(adUserPortKey int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func AllSlavesActiveOption(allSlavesActive int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func ArpIntervalOption(arpInterval int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func ArpIPTargetsOption(arpIPTargets []net.IP) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func ArpValidateOption(arpValidate int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func ArpAllTargetsOption(arpAllTargets int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func DownDelayOption(downDelay int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func FailOverMacOption(failOverMac int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func LacpRateOption(lacpRate int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func MiimonOption(miimon int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func MinLinksOption(minLinks int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func PacketsPerSlaveOption(packetsPerSlave int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func PrimaryOption(primary int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func PrimaryReselectOption(primaryReselect int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func TlbDynamicLbOption(tlbDynamicLb int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func UpDelayOption(upDelay int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func UseCarrierOption(useCarrier int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func XmitHashPolicyOption(xmitHashPolicy int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func LpIntervalOption(lpInterval int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func ResendIgmpOption(resendIgmp int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func NumPeerNotifOption(numPeerNotif int) BondOptionFunc {
	_ = "STUB: not implemented"
	return *new(BondOptionFunc)
}

func GetAllIntBondOptions(bondOptions *BondOptions, bondOptionFuncs []BondOptionFunc) []BondOptionFunc {
	_ = "STUB: not implemented"
	return nil
}

func getVlanIfaceName(master string, vlanID int) string { _ = "STUB: not implemented"; return "" }

func checkInterfaceWithSameVlan(vlanID int, vlanInterface string) error {
	_ = "STUB: not implemented"
	return nil
}
