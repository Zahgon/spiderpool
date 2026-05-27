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
)

func GenerateExampleReplicaSetYaml(rsName, namespace string, replica int32) *appsv1.ReplicaSet {
	_ = "STUB: not implemented"
	return nil
}

func ScaleReplicasetUntilExpectedReplicas(ctx context.Context, frame *e2e.Framework, rs *appsv1.ReplicaSet, expectedReplicas int, scalePodRun bool) (addedPod, removedPod []corev1.Pod, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// return the diff pod
