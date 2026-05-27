// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	e2e "github.com/spidernet-io/e2eframework/framework"
	corev1 "k8s.io/api/core/v1"
)

func SetNamespaceIppoolAnnotation(IppoolAnnoValue []string, nsObject *corev1.Namespace, PoolNameList []string, keyAnno string) {
	_ = "STUB: not implemented"
	return
}

func GeneratePodIPPoolAnnotations(frame *e2e.Framework, _ string, v4PoolNameList, v6PoolNameList []string) string {
	_ = "STUB: not implemented"
	return ""
}

func GeneratePodIPPoolsAnnotations(frame *e2e.Framework, nic string, cleanGateway bool, v4PoolNameList, v6PoolNameList []string) string {
	_ = "STUB: not implemented"
	return ""
}
