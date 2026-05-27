// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package gcmanager

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	ktypes "k8s.io/apimachinery/pkg/types"

	"github.com/spidernet-io/spiderpool/pkg/lock"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

type PodDBer interface {
	DeletePodEntry(namespace, podName string)
	ApplyPodEntry(podEntry *PodEntry) error
	ListAllPodEntries() []PodEntry
}

// PodEntry represents a pod cache
type PodEntry struct {
	PodName   string
	Namespace string
	NodeName  string
	UID       string

	EntryUpdateTime     time.Time
	TracingStartTime    time.Time
	TracingGracefulTime time.Duration
	TracingStopTime     time.Time

	PodTracingReason types.PodStatus
	NumRequeues      int
}

// PodDatabase represents controller PodEntry database
type PodDatabase struct {
	lock.RWMutex
	pods   map[ktypes.NamespacedName]PodEntry
	maxCap int
}

func NewPodDBer(maxDatabaseCap int) PodDBer { _ = "STUB: not implemented"; return *new(PodDBer) }

func (p *PodDatabase) DeletePodEntry(namespace, podName string) { _ = "STUB: not implemented"; return }

// already deleted

func (p *PodDatabase) ListAllPodEntries() []PodEntry { _ = "STUB: not implemented"; return nil }

func (p *PodDatabase) ApplyPodEntry(podEntry *PodEntry) error {
	_ = "STUB: not implemented"
	return nil
}

// diff and fresh the DB

// buildPodEntry will build PodEntry with the given args, it serves for Pod Informer event hooks and scanAll
// for Pod Informer event hooks, if the podEntry is nil, we don't tracing it
// for scanAll, if the podEntry is nil, we will not GC it's IP
func (s *SpiderGC) buildPodEntry(oldPod, currentPod *corev1.Pod, deleted bool) (*PodEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check StatefulSet pod, we will trace it if its controller StatefulSet object was deleted or decreased
// its replicas and the pod index was out of the replicas.

// check kubevirt vm pod, we will trace it if its controller is no longer exist

// deleted pod

// stop time

// case: current pod is 'Terminating', no old pod

// case: current pod is 'Terminating', old pod wasn't 'Terminating' phase

// case: both of current pod and old pod are 'Terminating', but 'TerminationGracePeriodSeconds' changed

// case: current pod is 'Succeeded|Failed', no old pod

// case: current pod is 'Succeeded|Failed', old pod wasn't 'Succeeded|Failed' phase

// case: both of current pod and old pod are 'Terminating', but 'TerminationGracePeriodSeconds' changed

// Running | Unknown

// check terminating Pod corresponding Node status

// grace period

// stop time

// start time

// grace period

// stop time

// computeSucceededOrFailedPodTerminatingTime will compute terminating start time, stop time and graceful period for 'Succeeded | Failed' phase pod
func (s *SpiderGC) computeSucceededOrFailedPodTerminatingTime(podYaml *corev1.Pod) (terminatingStartTime, terminatingStopTime time.Time, gracefulTime time.Duration, err error) {
	_ = "STUB: not implemented"
	// check container numbers
	return *new(time.Time), *new(time.Time), *new(time.Duration), nil
}

// compute Succeeded | Failed pod start time

// graceful period

// stop time

func (s *SpiderGC) isValidStatefulSetPod(ctx context.Context, currentPod *corev1.Pod) (isValidStsPod bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check StatefulSet pod, we will trace it if its controller StatefulSet object was deleted or decreased its replicas and the pod index was out of the replicas.

// StatefulSet pod restarted, no need to trace it.

func (s *SpiderGC) isValidKubevirtVMIPod(ctx context.Context, currentPod *corev1.Pod) (isKubevirtVMIPod bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// check StatefulSet pod, we will trace it if its controller StatefulSet object was deleted or decreased its replicas and the pod index was out of the replicas.
