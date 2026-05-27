// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package podmanager

import (
	"context"

	crdclientset "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/spiderpool/pkg/types"
)

type PodManager interface {
	GetPodByName(ctx context.Context, namespace, podName string, cached bool) (*corev1.Pod, error)
	ListPods(ctx context.Context, cached bool, opts ...client.ListOption) (*corev1.PodList, error)
	GetPodTopController(ctx context.Context, pod *corev1.Pod) (types.PodTopController, error)
}

type podManager struct {
	client       client.Client
	apiReader    client.Reader
	SpiderClient crdclientset.Interface
}

func NewPodManager(client client.Client, apiReader client.Reader) (PodManager, error) {
	_ = "STUB: not implemented"
	return *new(PodManager), nil
}

func (pm *podManager) GetPodByName(ctx context.Context, namespace, podName string, cached bool) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pm *podManager) ListPods(ctx context.Context, cached bool, opts ...client.ListOption) (*corev1.PodList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPodTopController will find the pod top owner controller with the given pod.
// For example, once we create a deployment then it will create replicaset and the replicaset will create pods.
// So, the pods' top owner is deployment. That's what the method implements.
// Notice: if the application is a third party controller, the types.PodTopController property App would be nil!
func (pm *podManager) GetPodTopController(ctx context.Context, pod *corev1.Pod) (types.PodTopController, error) {
	_ = "STUB: not implemented"
	return *new(types.PodTopController), nil
}

// pod.APIVersion is empty string

// third party controller

// deployment.APIVersion is empty string

// daemonSet.APIVersion is empty string

// statefulSet.APIVersion is empty string
