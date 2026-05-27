// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/spidernet-io/spiderpool/api/v1/controller/server/restapi/runtime"
)

// Singleton
var (
	httpGetControllerStartup   = &_httpGetControllerStartup{controllerContext}
	httpGetControllerReadiness = &_httpGetControllerReadiness{controllerContext}
	httpGetControllerLiveness  = &_httpGetControllerLiveness{controllerContext}
)

type _httpGetControllerStartup struct {
	*ControllerContext
}

// Handle handles GET requests for k8s startup probe.
func (g *_httpGetControllerStartup) Handle(params runtime.GetRuntimeStartupParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

type _httpGetControllerReadiness struct {
	*ControllerContext
}

// Handle handles GET requests for k8s readiness probe.
func (g *_httpGetControllerReadiness) Handle(params runtime.GetRuntimeReadinessParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}

type _httpGetControllerLiveness struct {
	*ControllerContext
}

// Handle handles GET requests for k8s liveness probe.
func (g *_httpGetControllerLiveness) Handle(params runtime.GetRuntimeLivenessParams) middleware.Responder {
	_ = "STUB: not implemented"
	return *new(middleware.Responder)
}
