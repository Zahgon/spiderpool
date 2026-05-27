// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/spidernet-io/spiderpool/api/v1/agent/server/restapi/runtime"
)

// Singleton
var (
	httpGetAgentStartup   = &_httpGetAgentStartup{agentContext}
	httpGetAgentReadiness = &_httpGetAgentReadiness{agentContext}
	httpGetAgentLiveness  = &_httpGetAgentLiveness{agentContext}
)

type _httpGetAgentStartup struct {
	*AgentContext
}

// Handle handles GET requests for k8s startup probe.
func (g *_httpGetAgentStartup) Handle(params runtime.GetRuntimeStartupParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

type _httpGetAgentReadiness struct {
	*AgentContext
}

// Handle handles GET requests for k8s readiness probe.
func (g *_httpGetAgentReadiness) Handle(params runtime.GetRuntimeReadinessParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

type _httpGetAgentLiveness struct {
	*AgentContext
}

// Handle handles GET requests for k8s liveness probe.
func (g *_httpGetAgentLiveness) Handle(params runtime.GetRuntimeLivenessParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}
