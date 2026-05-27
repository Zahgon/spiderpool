// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package coordinatormanager

import (
	"context"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var WebhookLogger *zap.Logger

type CoordinatorWebhook struct{}

func (cw *CoordinatorWebhook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomDefaulter = (*CoordinatorWebhook)(nil)

// Default implements webhook.CustomDefaulter so a webhook will be registered for the type.
func (cw *CoordinatorWebhook) Default(ctx context.Context, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomValidator = (*CoordinatorWebhook)(nil)

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type.
func (cw *CoordinatorWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type.
func (cw *CoordinatorWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type.
func (cw *CoordinatorWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
