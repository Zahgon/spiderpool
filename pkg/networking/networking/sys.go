// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package networking

const (
	SysClassNetDevicePath = "/sys/class/net"
	SysBusPciDevicesPath  = "/sys/bus/pci/devices"
)

func GetPfNameFromVfDeviceID(vfDeviceID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetPfDeviceIDFromVF(vfDeviceID string) (string, error) {
	_ = "STUB: not implemented"
	// First try the traditional approach via sysfs (works in host namespace)
	return "", nil
}

// Read the path that the symlink points to

func GetPfNameFromPfDeviceID(pfDeviceID string) (string, error) {
	_ = "STUB: not implemented"
	// Get the network interface name from PCI address
	return "", nil
}
