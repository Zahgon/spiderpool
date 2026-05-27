// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package nodemanager

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NodeManager interface {
	GetNodeByName(ctx context.Context, nodeName string, cached bool) (*corev1.Node, error)
	ListNodes(ctx context.Context, cached bool, opts ...client.ListOption) (*corev1.NodeList, error)
}

type nodeManager struct {
	client    client.Client
	apiReader client.Reader
}

func NewNodeManager(client client.Client, apiReader client.Reader) (NodeManager, error) {
	_ = "STUB: not implemented"
	return *new(NodeManager), nil
}

func (nm *nodeManager) GetNodeByName(ctx context.Context, nodeName string, cached bool) (*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (nm *nodeManager) ListNodes(ctx context.Context, cached bool, opts ...client.ListOption) (*corev1.NodeList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
