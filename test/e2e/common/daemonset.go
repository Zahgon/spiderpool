// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Package common provides test utilities for E2E tests.
package common

import (
	//nolint:all // Standard for Ginkgo/Gomega BDD testing framework
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
)

func GenerateExampleDaemonSetYaml(dsName, namespace string) *appsv1.DaemonSet {
	_ = "STUB: not implemented"
	return nil
}
