// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package gcmanager

import (
	"context"
	"fmt"
)

var errRequeue = fmt.Errorf("requeue")

// tracePodWorker will circle traverse PodEntry database
func (s *SpiderGC) tracePodWorker(ctx context.Context) { _ = "STUB: not implemented"; return }

// handlePodEntryForTracingTimeOut check the given podEntry whether out of time. If so, just send a signal to execute gc
func (s *SpiderGC) handlePodEntryForTracingTimeOut(podEntry *PodEntry) {
	_ = "STUB: not implemented"
	return
}

// If the statefulset application quickly experiences scaling down and up,
// check whether `Status.PodIPs` is empty to determine whether the Pod in the current K8S has completed the normal IP release to avoid releasing the wrong IP.

// the pod will be handled next time.

// not time out

// releaseIPPoolIPExecutor receive signals to execute gc IP
func (s *SpiderGC) releaseIPPoolIPExecutor(ctx context.Context, workerIndex int) {
	_ = "STUB: not implemented"
	return
}

// Pod has the same name as SpiderEndpoint, but the UID does not match.
// Such SpiderEndpoint should be reclaim, but because the IPPool name used by SpiderEndpoint cannot be tracked,
// it will be reclaim later via GC All

// we need to gather the pod corresponding SpiderEndpoint allocation data to get the used history IPs.

// release pod used history IPs

// delete StatefulSet/kubevirtVMI wep (other controller wep has OwnerReference, its lifecycle is same with pod)
