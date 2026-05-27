// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	agentOpenAPIServer "github.com/spidernet-io/spiderpool/api/v1/agent/server"
)

// newAgentOpenAPIHttpServer instantiates a new instance of the agent OpenAPI server on the http.
func newAgentOpenAPIHttpServer() (*agentOpenAPIServer.Server, error) {
	_ = "STUB: not implemented"
	// read yaml spec
	return nil, nil
}

// create new service API

// set spiderpool logger as api logger

// runtime API

// new agent OpenAPI server with api

// spiderpool-agent component owns Unix server and Http server, the Unix server uses for IPAM plugin interaction,
// and the Http server uses for K8s or CLI command.
// In spider-agent openapi.yaml, we already set x-schemes with value 'unix', so we need set Http server's listener with value 'http'.

// configure API and handlers with some default values.
