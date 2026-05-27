// This file is safe to edit. Once it exists it will not be overwritten

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"crypto/tls"
	"net/http"

	"github.com/spidernet-io/spiderpool/api/v1/agent/server/restapi"
)

//go:generate swagger generate server --target ../../agent --name SpiderpoolAgentAPI --spec ../openapi.yaml --api-package restapi --server-package server --principal interface{} --default-scheme unix --exclude-main

func configureFlags(api *restapi.SpiderpoolAgentAPIAPI) {
	_ = "STUB: not implemented"
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
	return
}

func configureAPI(api *restapi.SpiderpoolAgentAPIAPI) http.Handler {
	_ = "STUB: not implemented"
	// configure the api here
	return *new(http.Handler)
}

// Set your custom logger if needed. Default one is log.Printf
// Expected interface func(string, ...interface{})
//
// Example:
// api.Logger = log.Printf

// To continue using redoc as your UI, uncomment the following line
// api.UseRedoc()

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	_ = "STUB: not implemented"
	// Make all necessary changes to the TLS configuration here.
	return
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	_ = "STUB: not implemented"

	// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation.
	return
}

func setupMiddlewares(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"

	// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
	// So this is a good place to plug in a panic handling middleware, logging and metrics.
	return *new(http.Handler)
}

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
