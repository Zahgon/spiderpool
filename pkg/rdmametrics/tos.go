// Copyright 2025 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package rdmametrics

import (
	"os"
)

type DeviceTrafficClass struct {
	NetDevName   string
	IfName       string
	TrafficClass int
}

const prefix = "Global tclass="

var (
	statFunc     = os.Stat
	readFileFunc = os.ReadFile
)

func GetDeviceTrafficClass(impl NetlinkImpl) ([]DeviceTrafficClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
