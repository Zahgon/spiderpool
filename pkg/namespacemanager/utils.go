// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package namespacemanager

import (
	corev1 "k8s.io/api/core/v1"
)

func GetNSDefaultPools(ns *corev1.Namespace) ([]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
