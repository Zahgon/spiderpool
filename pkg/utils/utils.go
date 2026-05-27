// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	corev1 "k8s.io/api/core/v1"
)

// GetDefaultCNIConfPath according to the provided CNI file path (default is /etc/cni/net.d),
// return the first CNI configuration file path under this path.
func GetDefaultCNIConfPath(cniDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetDefaultCniName according to the provided CNI file path (default is /etc/cni/net.d),
// the first CNI configuration file under this path is parsed and its name is returned
func GetDefaultCniName(cniDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func findDefaultCNIConf(cniDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func fetchCniNameFromPath(cniPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ExtractK8sCIDRFromKubeadmConfigMap(cm *corev1.ConfigMap) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ExtractK8sCIDRFromKCMPod(kcm *corev1.Pod) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}
