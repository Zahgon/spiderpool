// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package statefulsetmanager

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type StatefulSetManager interface {
	GetStatefulSetByName(ctx context.Context, namespace, name string, cached bool) (*appsv1.StatefulSet, error)
	ListStatefulSets(ctx context.Context, cached bool, opts ...client.ListOption) (*appsv1.StatefulSetList, error)
	IsValidStatefulSetPod(ctx context.Context, namespace, podName, podControllerType string) (bool, error)
}

type statefulSetManager struct {
	client    client.Client
	apiReader client.Reader
}

func NewStatefulSetManager(client client.Client, apiReader client.Reader) (StatefulSetManager, error) {
	_ = "STUB: not implemented"
	return *new(StatefulSetManager), nil
}

func (sm *statefulSetManager) GetStatefulSetByName(ctx context.Context, namespace, name string, cached bool) (*appsv1.StatefulSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sm *statefulSetManager) ListStatefulSets(ctx context.Context, cached bool, opts ...client.ListOption) (*appsv1.StatefulSetList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsValidStatefulSetPod only serves for StatefulSet pod, it will check the pod whether need to be cleaned up with the given params podNS, podName.
// Once the pod's controller StatefulSet was deleted, the pod's corresponding IPPool IP and Endpoint need to be cleaned up.
// Or the pod's controller StatefulSet decreased its replicas and the pod's index is out of replicas, it needs to be cleaned up too.
func (sm *statefulSetManager) IsValidStatefulSetPod(ctx context.Context, namespace, podName, podControllerType string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Ref: https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#start-ordinal

// The Pod controlled by StatefulSet is created or re-created.

// StatefulSet scaled down.
