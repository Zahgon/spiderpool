// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	e2e "github.com/spidernet-io/e2eframework/framework"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GenerateExampleStatefulSetYaml(stsName, namespace string, replica int32) *appsv1.StatefulSet {
	_ = "STUB: not implemented"
	return nil
}

func ScaleStatefulsetUntilExpectedReplicas(ctx context.Context, frame *e2e.Framework, sts *appsv1.StatefulSet, expectedReplicas int, scalePodRun bool) (addedPod, removedPod []corev1.Pod, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// return the diff pod

func PatchStatefulSet(frame *e2e.Framework, desiredStatefulSet, originalStatefulSet *appsv1.StatefulSet, opts ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func RestartAndValidateStatefulSetPodIP(frame *e2e.Framework, label map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func recordStatefulSetPodIP(podList *corev1.PodList) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
