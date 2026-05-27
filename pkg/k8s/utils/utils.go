// Copyright 2024 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func DeleteWebhookConfiguration(ctx context.Context, c client.Client, name string, obj client.Object) error {
	_ = "STUB: not implemented"
	return nil
}
