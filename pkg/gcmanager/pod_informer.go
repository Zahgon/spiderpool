// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package gcmanager

import (
	"context"
)

// startPodInformer will set up k8s pod informer in circle
func (s *SpiderGC) startPodInformer(ctx context.Context) { _ = "STUB: not implemented"; return }

// Let the leader trigger IP GC scan all.
// When the spiderpool-controller restarted, it will trigger IP GC scan all first.
// If the pod informer not starts and the user delete some pods, this will lead to IP leakage.

// onPodAdd represents Pod informer Add Event
func (s *SpiderGC) onPodAdd(obj interface{}) {
	_ = "STUB: not implemented"
	// backup controller could be elected as master
	return
}

// flush the pod database

// onPodUpdate represents Pod informer Update Event
func (s *SpiderGC) onPodUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	// backup controller could be elected as master
	return
}

// flush the pod database

// onPodDel represents Pod informer Delete Event
func (s *SpiderGC) onPodDel(obj interface{}) {
	_ = "STUB: not implemented"
	// backup controller could be elected as master
	return
}
