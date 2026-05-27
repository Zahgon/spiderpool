// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package podmanager

import (
	"context"

	"github.com/spidernet-io/spiderpool/pkg/constant"
	crdclientset "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	"github.com/spidernet-io/spiderpool/pkg/namespacemanager"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var PodWebhookExcludeNamespaces = []string{
	metav1.NamespaceSystem,
	metav1.NamespacePublic,
	constant.Spiderpool,
	"metallb-system",
	"istio-system",
	// more system namespaces to be added
}

type PodWebhook interface {
	admission.CustomDefaulter
	admission.CustomValidator
}

type PWebhook struct {
	spiderClient crdclientset.Interface
	nsManager    namespacemanager.NamespaceManager
}

// InitPodWebhook initializes the pod webhook.
// It sets up the mutating webhook for pods and registers it with the manager.
// Parameters:
//   - mgr: The controller manager
//
// Returns an error if initialization fails.
func InitPodWebhook(mgr ctrl.Manager, nsManager namespacemanager.NamespaceManager) error {
	_ = "STUB: not implemented"
	return nil
}

// setup mutating webhook for pods

// Default implements the defaulting webhook for pods.
// It injects network resources into the pod if it has the appropriate annotation.
// Parameters:
//   - ctx: The context
//   - obj: The runtime object (expected to be a Pod)
//
// Returns an error if defaulting fails.
func (pw *PWebhook) Default(ctx context.Context, obj runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateCreate implements the validation webhook for pod creation.
// Currently, it performs no validation and always returns nil.
func (pw *PWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"

	// ValidateUpdate implements the validation webhook for pod updates.
	// Currently, it performs no validation and always returns nil.
	return *new(admission.Warnings), nil
}

func (pw *PWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"

	// ValidateDelete implements the validation webhook for pod deletion.
	// Currently, it performs no validation and always returns nil.
	return *new(admission.Warnings), nil
}

func (pw *PWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	_ = "STUB: not implemented"
	return *new(admission.Warnings), nil
}
