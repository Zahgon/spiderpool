// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package subnetmanager

import (
	"context"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var WebhookLogger *zap.Logger

type SubnetWebhook struct {
	Client    client.Client
	APIReader client.Reader

	EnableIPv4                              bool
	EnableIPv6                              bool
	EnableValidatingResourcesDeletedWebhook bool
}

func (sw *SubnetWebhook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomDefaulter = (*SubnetWebhook)(nil)

// Default implements webhook.CustomDefaulter so a webhook will be registered for the type.
func (sw *SubnetWebhook) Default(ctx context.Context, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

var _ webhook.CustomValidator = (*SubnetWebhook)(nil)

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type.
func (sw *SubnetWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// the user will receive the following errors rather than K8S API server specific typed errors.
// Refer to https://github.com/spidernet-io/spiderpool/issues/3321

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type.
func (sw *SubnetWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type.
func (sw *SubnetWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
