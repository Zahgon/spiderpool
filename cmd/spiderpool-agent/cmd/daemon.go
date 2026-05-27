// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"os"
)

// DaemonMain runs agentContext handlers.
func DaemonMain() {
	_ = "STUB: not implemented"
	// Set logger level and re-init global logger.
	return
}

// Print version info for debug.

// Set golang max procs.

// Load spiderpool's global Comfigmap.

// setup sysctls

// Set up gops.

// Set up pyroscope.

// These 2 lines are only required if you're using mutex or block profiling

// additional

// init managers...

// clean up unix socket path legacy, it won't return an error if it doesn't exist

// WatchSignal notifies the signal to shut down agentContext handlers.
func WatchSignal(sigCh chan os.Signal) { _ = "STUB: not implemented"; return }

// Cancel the internal context of spiderpool-agent.
// This stops things like the runtime manager, GC, etc.

// shut down agent http server

// shut down agent unix server

// others...

func waitAPIServerReady(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// This client does not query any Kubernetes resources, it is only used
// to detect whether the API Server's readiness probe is ready, so there
// is no need to add any decoder.

// Request API Server every 2 seconds until API Server is ready or all 15
// retries have timed out. (total cost: 2 * 15 = 30s)

func initAgentServiceManagers(ctx context.Context) { _ = "STUB: not implemented"; return }

// sysctlConfig set default sysctl configs,Notice: ignore not exist sysctl configs as
// possible.
func sysctlConfig(enableIPv4, enableIPv6 bool) error {
	_ = "STUB: not implemented"
	// setup default sysctl config
	return nil
}
