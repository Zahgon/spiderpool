// Copyright 2024 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package openapi

import (
	"net/http"
)

// NewWebhookHealthCheckClient creates one http client which serves for webhook health check
func NewWebhookHealthCheckClient() *http.Client { _ = "STUB: not implemented"; return nil }

// WebhookHealthyCheck servers for spiderpool controller readiness and liveness probe.
// This is a Layer7 check.
func WebhookHealthyCheck(httpClient *http.Client, webhookPort string, url *string) error {
	_ = "STUB: not implemented"
	return nil
}
