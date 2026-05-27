// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Package common provides test utilities for E2E tests.
package common

import (
	"context"

	//nolint:all // Standard for Ginkgo/Gomega BDD testing framework
	. "github.com/onsi/ginkgo/v2"
	//nolint:all // Standard for Ginkgo/Gomega BDD testing framework
	. "github.com/onsi/gomega"
	e2e "github.com/spidernet-io/e2eframework/framework"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func GenerateExampleDeploymentYaml(dpmName, namespace string, replica int32) *appsv1.Deployment {
	_ = "STUB: not implemented"
	return nil
}

func GenerateDraDeploymentYaml(dpmName, claim, namespace string, replica int32) *appsv1.Deployment {
	_ = "STUB: not implemented"
	return nil
}

func ScaleDeployUntilExpectedReplicas(ctx context.Context, frame *e2e.Framework, deploy *appsv1.Deployment, expectedReplicas int, scalePodRun bool) (addedPod, removedPod []corev1.Pod, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// return the diff pod

func CreateDeployUntilExpectedReplicas(frame *e2e.Framework, deploy *appsv1.Deployment, ctx context.Context) (pods *corev1.PodList, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create Deployment with types.AnnoPodIPPoolValue
func CreateDeployWithPodAnnoation(frame *e2e.Framework, name, namespace string, deployOriginialNum int, nic string, v4PoolNameList, v6PoolNameList []string) (deploy *appsv1.Deployment) {
	_ = "STUB: not implemented"
	return nil
}

// Create Deployment until the ip assignment is successful
func CreateDeployUnitlReadyCheckInIppool(frame *e2e.Framework, depName, namespaceName string, podNum int32, v4PoolNameList, v6PoolNameList []string) {
	_ = "STUB: not implemented"
	return
}

// get pod list

// check pod ip record still in this ippool

func BatchCreateDeploymentUntilReady(ctx context.Context, frame *e2e.Framework, expectedNum, replicas int, namespace string, annotationMap map[string]string) []string {
	_ = "STUB: not implemented"
	return nil
}
