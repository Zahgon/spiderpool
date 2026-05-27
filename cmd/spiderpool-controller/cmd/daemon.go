// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"os"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// DaemonMain runs controllerContext handlers.
func DaemonMain() {
	_ = "STUB: not implemented"
	// Set logger level and re-init global logger.
	return
}

// Print version info for debug.

// Set golang max procs.

// Load spiderpool's global Comfigmap.

// Set up gops.

// Set up pyroscope.

// These 2 lines are only required if you're using mutex or block profiling

// additional

// init managers...

// The CRD webhook of Spiderpool must be started before informer, so that
// informer can normally request to some CRs in the cluster without being
// disturbed by an abnormal webhook.

// WatchSignal notifies the signal to shut down controllerContext handlers.
func WatchSignal(sigCh chan os.Signal) { _ = "STUB: not implemented"; return }

// Cancel the internal context of spiderpool-controller.
// This stops things like the runtime manager, GC, etc.

// shut down http server

// others...

func initControllerServiceManagers(ctx context.Context) { _ = "STUB: not implemented"; return }

func initGCManager(ctx context.Context) {
	_ = "STUB: not implemented"
	// EnableStatefulSet was determined by Configmap.
	return
}

// EnableKubevirtStaticIP was determined by Configmap.

// EnableCleanOutdatedEndpoint was determined by Configmap.

func initSpiderControllerLeaderElect(ctx context.Context) { _ = "STUB: not implemented"; return }

// initK8sClientSet will new kubernetes Clientset
func initK8sClientSet() (*kubernetes.Clientset, error) { _ = "STUB: not implemented"; return nil, nil }

func initDynamicClient() (*dynamic.DynamicClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setupInformers will run IPPool,Subnet... informers,
// because these informers count on webhook
func setupInformers(k8sClient *kubernetes.Clientset) {
	_ = "STUB: not implemented"
	// start SpiderIPPool informer
	return
}

func checkWebhookReady() { _ = "STUB: not implemented"; return }
