// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	"context"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	e2e "github.com/spidernet-io/e2eframework/framework"
	"github.com/spidernet-io/spiderpool/pkg/types"
	corev1 "k8s.io/api/core/v1"
)

func GenerateExamplePodYaml(podName, namespace string) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func CreatePodUntilReady(frame *e2e.Framework, podYaml *corev1.Pod, podName, namespace string, waitPodStartTimeout time.Duration) (pod *corev1.Pod, podIPv4, podIPv6 string) {
	_ = "STUB: not implemented"
	// create pod
	return nil, "", ""
}

// wait for pod ip

func CreatePodWithAnnoPodIPPool(frame *e2e.Framework, podName, namespace string, annoPodIPPoolValue types.AnnoPodIPPoolValue) {
	_ = "STUB: not implemented"
	return
}

func CheckPodIPReadyByLabel(frame *e2e.Framework, label map[string]string, v4PoolNameList, v6PoolNameList []string) *corev1.PodList {
	_ = "STUB: not implemented"
	return nil
}

// Get the rebuild pod list

// Succeeded to assign ipv4、ipv6 ip for pod

// check pod ip recorded in ippool

func DeletePods(frame *e2e.Framework, opts ...client.DeleteAllOfOption) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidatePodIPConflict(podList *corev1.PodList) error { _ = "STUB: not implemented"; return nil }

func GetPodNetworkInfo(ctx context.Context, frame *e2e.Framework, podList *corev1.PodList) error {
	_ = "STUB: not implemented"
	return nil
}
