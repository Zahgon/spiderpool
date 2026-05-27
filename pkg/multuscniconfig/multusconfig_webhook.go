// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package multuscniconfig

import (
	"context"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var logger *zap.Logger

type MultusConfigWebhook struct {
	APIReader client.Reader
}

func (mcw *MultusConfigWebhook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomDefaulter = (*MultusConfigWebhook)(nil)

// Default implements admission.CustomDefaulter.
func (*MultusConfigWebhook) Default(ctx context.Context, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomValidator = (*MultusConfigWebhook)(nil)

func (mcw *MultusConfigWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

func (mcw *MultusConfigWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateDelete will implement something just like kubernetes Foreground cascade deletion to delete the MultusConfig corresponding net-attach-def firstly
// Since the MultusConf doesn't have Finalizer, you could delete it as soon as possible and we can't filter it to delete the net-attach-def at first.
func (mcw *MultusConfigWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
