// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

/**
* Copyright (c) 2017 Intel Corporation
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package multuscniconfig

import (
	netv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"

	coordinatorcmd "github.com/spidernet-io/spiderpool/cmd/coordinator/cmd"
	spiderpoolcmd "github.com/spidernet-io/spiderpool/cmd/spiderpool/cmd"
	v2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

type MacvlanNetConf struct {
	Type   string                    `json:"type"`
	Master string                    `json:"master"`
	Mode   string                    `json:"mode"`
	MTU    *int32                    `json:"mtu,omitempty"`
	IPAM   *spiderpoolcmd.IPAMConfig `json:"ipam,omitempty"`
}

type IPvlanNetConf struct {
	Type   string                    `json:"type"`
	Master string                    `json:"master"`
	MTU    *int32                    `json:"mtu,omitempty"`
	IPAM   *spiderpoolcmd.IPAMConfig `json:"ipam,omitempty"`
}

type VlanNetConf struct {
	Type   string                    `json:"type"`
	Master string                    `json:"master"`
	VlanID int32                     `json:"vlanId"`
	MTU    *int32                    `json:"mtu,omitempty"`
	IPAM   *spiderpoolcmd.IPAMConfig `json:"ipam,omitempty"`
}

type SRIOVNetConf struct {
	Vlan *int32 `json:"vlan,omitempty"`
	// Mbps, 0 = disable rate limiting
	MinTxRate *int `json:"min_tx_rate,omitempty"`
	// Mbps, 0 = disable rate limiting
	MaxTxRate *int                      `json:"max_tx_rate,omitempty"`
	Type      string                    `json:"type"`
	DeviceID  string                    `json:"deviceID,omitempty"`
	IPAM      *spiderpoolcmd.IPAMConfig `json:"ipam,omitempty"`
}

type IBSRIOVNetConf struct {
	Type                string                    `json:"type"`
	Pkey                *string                   `json:"pkey,omitempty"`
	LinkState           *string                   `json:"link_state,omitempty"`
	RdmaIsolation       *bool                     `json:"rdmaIsolation,omitempty"`
	IBKubernetesEnabled *bool                     `json:"ibKubernetesEnabled,omitempty"`
	IPAM                *spiderpoolcmd.IPAMConfig `json:"ipam,omitempty"`
}

type IPoIBNetConf struct {
	Type   string                    `json:"type"`
	Master string                    `json:"master,omitempty"`
	IPAM   *spiderpoolcmd.IPAMConfig `json:"ipam,omitempty"`
}

type RdmaNetConf struct {
	Type string `json:"type"`
}

type OvsNetConf struct {
	Vlan     *int32                    `json:"vlan,omitempty"`
	Type     string                    `json:"type"`
	BrName   string                    `json:"bridge"`
	DeviceID string                    `json:"deviceID,omitempty"`
	IPAM     *spiderpoolcmd.IPAMConfig `json:"ipam,omitempty"`
	Trunk    []*v2beta1.Trunk          `json:"trunk,omitempty"`
}

type IfacerNetConf struct {
	VlanID     int                 `json:"vlanID,omitempty"`
	Type       string              `json:"type"`
	Interfaces []string            `json:"interfaces,omitempty"`
	Bond       *v2beta1.BondConfig `json:"bond,omitempty"`
}

type tuningConf struct {
	Type string `json:"type"`
	Mtu  int32  `json:"mtu,omitempty"`
}

type CoordinatorConfig struct {
	TxQueueLen         *int                `json:"txQueueLen,omitempty"`
	IPConflict         *bool               `json:"detectIPConflict,omitempty"`
	DetectGateway      *bool               `json:"detectGateway,omitempty"`
	VethLinkAddress    string              `json:"vethLinkAddress,omitempty"`
	TunePodRoutes      *bool               `json:"tunePodRoutes,omitempty"`
	MacPrefix          string              `json:"podMACPrefix,omitempty"`
	Mode               coordinatorcmd.Mode `json:"mode,omitempty"`
	Type               string              `json:"type"`
	PodDefaultRouteNIC string              `json:"podDefaultRouteNic,omitempty"`
	PodRPFilter        *int                `json:"podRPFilter,omitempty" `
	OverlayPodCIDR     []string            `json:"overlayPodCIDR,omitempty"`
	ServiceCIDR        []string            `json:"serviceCIDR,omitempty"`
	HijackCIDR         []string            `json:"hijackCIDR,omitempty"`
}

func ParsePodNetworkAnnotation(podNetworks, defaultNamespace string) ([]*netv1.NetworkSelectionElement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Match Multus / network-attachment-definition-client ParseNetworkAnnotation: any of `[`, `{`, `"`
// indicates JSON (including pretty-printed arrays that do not start with `[{"`).

// Comma-delimited list of network attachment object names

// Remove leading and trailing whitespace.

// Parse network name (i.e. <namespace>/<network name>@<ifname>)

// validate MAC address

// validate GUID address

// validate IP address

// compatibility pre v3.2, will be removed in v4.0
// if n.DeprecatedInterfaceRequest != "" && n.InterfaceRequest == "" {
//	n.InterfaceRequest = n.DeprecatedInterfaceRequest
// }

func ParsePodNetworkObjectName(podnetwork string) (string, string, string, error) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

// Check and see if each item matches the specification for valid attachment name.
// "Valid attachment names must be comprised of units of the DNS-1123 label format"
// [a-z0-9]([-a-z0-9]*[a-z0-9])?
// It must start and end alphanumerically.

// Validate interface name: must be shorter than IFNAMSIZ (typically 16) and
// must not contain spaces or forward slashes, matching upstream multus-cni.

// ResourceName returns the appropriate resource name based on the CNI type and configuration
// of the given SpiderMultusConfig.
func ResourceName(smc *v2beta1.SpiderMultusConfig) string { _ = "STUB: not implemented"; return "" }

// For Macvlan CNI, return RDMA resource name if RDMA is enabled

func ValidateRdmaResouce(name, namespace, rdmaResourceName string, ippools *v2beta1.SpiderpoolPools) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateNetworkResouce(name, namespace, resourceName string, ippools *v2beta1.SpiderpoolPools) error {
	_ = "STUB: not implemented"
	return nil
}
