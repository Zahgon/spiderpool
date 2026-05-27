// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package reservedipmanager

import (
	"context"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var WebhookLogger *zap.Logger

type ReservedIPWebhook struct {
	EnableIPv4 bool
	EnableIPv6 bool
}

func (rw *ReservedIPWebhook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomDefaulter = (*ReservedIPWebhook)(nil)

// Default implements webhook.CustomDefaulter so a webhook will be registered for the type.
func (rw *ReservedIPWebhook) Default(ctx context.Context, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomValidator = (*ReservedIPWebhook)(nil)

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type.
func (rw *ReservedIPWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type.
func (rw *ReservedIPWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type.
func (rw *ReservedIPWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
