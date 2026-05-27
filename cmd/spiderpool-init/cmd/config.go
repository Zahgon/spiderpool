// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

const (
	ENVNamespace                = "SPIDERPOOL_NAMESPACE"
	ENVSpiderpoolControllerName = "SPIDERPOOL_CONTROLLER_NAME"
	ENVSpiderpoolAgentName      = "SPIDERPOOL_AGENT_NAME"

	ENVDefaultCoordinatorName             = "SPIDERPOOL_INIT_DEFAULT_COORDINATOR_NAME"
	ENVDefaultCoordinatorTuneMode         = "SPIDERPOOL_INIT_DEFAULT_COORDINATOR_MODE"
	ENVDefaultCoordinatorPodCIDRType      = "SPIDERPOOL_INIT_DEFAULT_COORDINATOR_POD_CIDR_TYPE"
	ENVDefaultCoordinatorDetectGateway    = "SPIDERPOOL_INIT_DEFAULT_COORDINATOR_DETECT_GATEWAY"
	ENVDefaultCoordinatorDetectIPConflict = "SPIDERPOOL_INIT_DEFAULT_COORDINATOR_DETECT_IP_CONFLICT"
	ENVDefaultCoordinatorTunePodRoutes    = "SPIDERPOOL_INIT_DEFAULT_COORDINATOR_TUNE_POD_ROUTES"
	ENVDefaultCoordiantorHijackCIDR       = "SPIDERPOOL_INIT_DEFAULT_COORDINATOR_HIJACK_CIDR"

	ENVDefaultIPv4SubnetName = "SPIDERPOOL_INIT_DEFAULT_IPV4_SUBNET_NAME"
	ENVDefaultIPv4IPPoolName = "SPIDERPOOL_INIT_DEFAULT_IPV4_IPPOOL_NAME"
	ENVDefaultIPv4CIDR       = "SPIDERPOOL_INIT_DEFAULT_IPV4_IPPOOL_SUBNET"
	ENVDefaultIPv4IPRanges   = "SPIDERPOOL_INIT_DEFAULT_IPV4_IPPOOL_IPRANGES"
	ENVDefaultIPv4Gateway    = "SPIDERPOOL_INIT_DEFAULT_IPV4_IPPOOL_GATEWAY"

	ENVDefaultIPv6SubnetName = "SPIDERPOOL_INIT_DEFAULT_IPV6_SUBNET_NAME"
	ENVDefaultIPv6IPPoolName = "SPIDERPOOL_INIT_DEFAULT_IPV6_IPPOOL_NAME"
	ENVDefaultIPv6CIDR       = "SPIDERPOOL_INIT_DEFAULT_IPV6_IPPOOL_SUBNET"
	ENVDefaultIPv6IPRanges   = "SPIDERPOOL_INIT_DEFAULT_IPV6_IPPOOL_IPRANGES"
	ENVDefaultIPv6Gateway    = "SPIDERPOOL_INIT_DEFAULT_IPV6_IPPOOL_GATEWAY"

	ENVEnableMultusConfig                = "SPIDERPOOL_INIT_ENABLE_MULTUS_CONFIG"
	ENVInstallMultusCNI                  = "SPIDERPOOL_INIT_INSTALL_MULTUS"
	ENVDefaultCNIDir                     = "SPIDERPOOL_INIT_DEFAULT_CNI_DIR"
	ENVDefaultCNIName                    = "SPIDERPOOL_INIT_DEFAULT_CNI_NAME"
	ENVDefaultCNINamespace               = "SPIDERPOOL_INIT_DEFAULT_CNI_NAMESPACE"
	ENVDefaultMultusConfigMap            = "SPIDERPOOL_INIT_MULTUS_CONFIGMAP"
	ENVDefaultReadinessFile              = "SPIDERPOOL_INIT_READINESS_FILE"
	ENVDefaultCoordinatorVethLinkAddress = "SPIDERPOOL_INIT_DEFAULT_COORDINATOR_VETH_LINK_ADDRESS"
)

var (
	legacyCalicoCniName = "k8s-pod-network"
	calicoCniName       = "calico"
	readinessFileName   = "/etc/spiderpool/ready"
)

type InitDefaultConfig struct {
	Namespace      string
	ControllerName string
	AgentName      string

	CoordinatorName               string
	CoordinatorMode               string
	CoordinatorPodCIDRType        string
	CoordinatorPodDefaultRouteNic string
	CoordinatorPodMACPrefix       string
	CoordinatorVethLinkAddress    string
	CoordinatorTunePodRoutes      bool
	CoordinatorHijackCIDR         []string

	V4SubnetName string
	V4IPPoolName string
	V4CIDR       string
	V4IPRanges   []string
	V4Gateway    string

	V6SubnetName string
	V6IPPoolName string
	V6CIDR       string
	V6IPRanges   []string
	V6Gateway    string

	// multuscniconfig
	enableMultusConfig  bool
	DefaultCNIDir       string
	DefaultCNIName      string
	DefaultCNINamespace string
	MultusConfigMap     string

	// readiness
	ReadinessFile string
}

func NewInitDefaultConfig() InitDefaultConfig {
	_ = "STUB: not implemented"
	return *new(InitDefaultConfig)
}

func parseENVAsDefault() InitDefaultConfig {
	_ = "STUB: not implemented"
	return *new(InitDefaultConfig)
}

// Coordinator

// IPv4

// IPv6

// parseCNIFromConfig parse cni's name and type from given cni config path
func parseCNIFromConfig(cniConfigPath string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func getMultusCniConfig(cniName, cniType string, ns string) *spiderpoolv2beta1.SpiderMultusConfig {
	_ = "STUB: not implemented"
	return nil
}

// change calico cni name from k8s-pod-network to calico
// more datails see:
// https://github.com/projectcalico/calico/issues/7837
