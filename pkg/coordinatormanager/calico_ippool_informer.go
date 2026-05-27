// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package coordinatormanager

import (
	"context"

	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

func NewCalicoIPPoolController(mgr ctrl.Manager, workQueue workqueue.RateLimitingInterface) (controller.Controller, error) {
	_ = "STUB: not implemented"
	return *new(controller.Controller), nil
}

type calicoIPPoolReconciler struct {
	client                     client.Client
	spiderCoordinatorWorkqueue workqueue.RateLimitingInterface
}

func (r *calicoIPPoolReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = "STUB: not implemented"
	return *new(ctrl.Result), nil
}
