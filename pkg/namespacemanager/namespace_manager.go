// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package namespacemanager

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NamespaceManager interface {
	GetNamespaceByName(ctx context.Context, nsName string, cached bool) (*corev1.Namespace, error)
	ListNamespaces(ctx context.Context, cached bool, opts ...client.ListOption) (*corev1.NamespaceList, error)
}

type namespaceManager struct {
	client    client.Client
	apiReader client.Reader
}

func NewNamespaceManager(client client.Client, apiReader client.Reader) (NamespaceManager, error) {
	_ = "STUB: not implemented"
	return *new(NamespaceManager), nil
}

func (nm *namespaceManager) GetNamespaceByName(ctx context.Context, nsName string, cached bool) (*corev1.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nm *namespaceManager) ListNamespaces(ctx context.Context, cached bool, opts ...client.ListOption) (*corev1.NamespaceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
