// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"os"
	"path/filepath"

	"github.com/containernetworking/cni/pkg/types"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
)

var (
	defaultLogPath          = "/var/log/spidernet/coordinator.log"
	defaultUnderlayVethName = "veth0"
	defaultMarkBit          = 0 // ox1
	// by default, k8s pod's first NIC is eth0
	defaultOverlayVethName  = "eth0"
	defaultPodRuleTable     = 100
	defaultHostRulePriority = 1000
	BinNamePlugin           = filepath.Base(os.Args[0])
)

type Mode string

const (
	ModeAuto     Mode = "auto"
	ModeUnderlay Mode = "underlay"
	ModeOverlay  Mode = "overlay"
	ModeDisable  Mode = "disable"
)

type Config struct {
	types.NetConf
	VethLinkAddress    string      `json:"vethLinkAddress,omitempty"`
	MacPrefix          string      `json:"podMACPrefix,omitempty"`
	MultusNicPrefix    string      `json:"multusNicPrefix,omitempty"`
	PodDefaultCniNic   string      `json:"podDefaultCniNic,omitempty"`
	OverlayPodCIDR     []string    `json:"overlayPodCIDR,omitempty"`
	ServiceCIDR        []string    `json:"serviceCIDR,omitempty"`
	HijackCIDR         []string    `json:"hijackCIDR,omitempty"`
	TunePodRoutes      *bool       `json:"tunePodRoutes,omitempty"`
	PodDefaultRouteNIC string      `json:"podDefaultRouteNic,omitempty"`
	Mode               Mode        `json:"mode,omitempty"`
	HostRuleTable      *int64      `json:"hostRuleTable,omitempty"`
	HostRPFilter       *int32      `json:"hostRPFilter,omitempty" `
	PodRPFilter        *int32      `json:"podRPFilter,omitempty" `
	TxQueueLen         *int64      `json:"txQueueLen,omitempty"`
	LogOptions         *LogOptions `json:"logOptions,omitempty"`
}

// DetectOptions enable ip conflicting check for pod's ip
type DetectOptions struct {
	Retry    int    `json:"retries,omitempty"`
	Interval string `json:"interval,omitempty"`
	TimeOut  string `json:"timeout,omitempty"`
}

type LogOptions struct {
	LogLevel        string `json:"logLevel"`
	LogFilePath     string `json:"logFile"`
	LogFileMaxSize  int    `json:"logMaxSize"`
	LogFileMaxAge   int    `json:"logMaxAge"`
	LogFileMaxCount int    `json:"logMaxCount"`
}

const (
	CniVersion030 = "0.3.0"
	CniVersion031 = "0.3.1"
	CniVersion040 = "0.4.0"
	CniVersion100 = "1.0.0"
)

// SupportCNIVersions indicate the CNI version that spiderpool support.
var SupportCNIVersions = []string{CniVersion030, CniVersion031, CniVersion040, CniVersion100}

// ParseConfig parses the supplied configuration (and prevResult) from stdin.
func ParseConfig(stdin []byte, coordinatorConfig *models.CoordinatorConfig) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// value must be -1,0/1/2

func validateHwPrefix(prefix string) error { _ = "STUB: not implemented"; return nil }

// prefix format like: 00:00、0a:1b

func ValidateRoutes(conf *Config, coordinatorConfig *models.CoordinatorConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRoutes(routes []string) error { _ = "STUB: not implemented"; return nil }

func validateRPFilterConfig(rpfilter *int32, coordinatorConfig int64) (*int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: negative number means disable
