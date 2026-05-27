// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package kubevirtmanager

import (
	"context"

	kubevirtv1 "kubevirt.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubevirtManager interface {
	IsValidVMPod(ctx context.Context, namespace, podControllerType, podControllerName string) (bool, error)
	GetVMByName(ctx context.Context, namespace, name string, cached bool) (*kubevirtv1.VirtualMachine, error)
	GetVMIByName(ctx context.Context, namespace, name string, cached bool) (*kubevirtv1.VirtualMachineInstance, error)
	GetVMIMByName(ctx context.Context, namespace, name string, cached bool) (*kubevirtv1.VirtualMachineInstanceMigration, error)
}

type kubevirtManager struct {
	client    client.Client
	apiReader client.Reader
}

func NewKubevirtManager(client client.Client, apiReader client.Reader) KubevirtManager {
	_ = "STUB: not implemented"
	return *new(KubevirtManager)
}

func (km *kubevirtManager) IsValidVMPod(ctx context.Context, namespace, podControllerType, podControllerName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// if the vmi was deleted, try to check its controller vm

// if the vmi is terminating and no owner controller vm, the pod is no longer valid.

// if the vmi is still alive, the pod is valid.

// is the vm is not exist, the pod is no longer valid

// is the vm is terminating, the pod is no longer valid

func (km *kubevirtManager) GetVMByName(ctx context.Context, namespace, name string, cached bool) (*kubevirtv1.VirtualMachine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (km *kubevirtManager) GetVMIByName(ctx context.Context, namespace, name string, cached bool) (*kubevirtv1.VirtualMachineInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (km *kubevirtManager) GetVMIMByName(ctx context.Context, namespace, name string, cached bool) (*kubevirtv1.VirtualMachineInstanceMigration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
