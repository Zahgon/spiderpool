// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package gcmanager

import (
	"context"

	"go.uber.org/zap"

	corev1 "k8s.io/api/core/v1"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

// monitorGCSignal will monitor signal from CLI, DefaultGCInterval
func (s *SpiderGC) monitorGCSignal(ctx context.Context) { _ = "STUB: not implemented"; return }

// In concurrency situation, the backup controller must execute scanAll

// The Elected controller will scan All with default GC interval

// CLI request

// discard the concurrent signal

// executeScanAll scans the whole pod and whole IPPoolList
func (s *SpiderGC) executeScanAll(ctx context.Context) { _ = "STUB: not implemented"; return }

// If we find the endpoint through the allocation information in the IP pool,
// it means that this endpoint is normal. Otherwise, it is very likely
// to be a leaked endpoint. We remove it from the suspiciousEndpointMap, so
// that it can be further cleaned up in the cleanOutdateEndpoint function.

// handle the pod not existed with the same name

// case: The pod in IPPool's ip-allocationDetail is not exist in k8s

// check should handle podIP via corresponding Node status and global gc flag

// check the pod status

// PodFailed means that all containers in the pod have terminated, and at least one container has
// terminated in a failure (exited with a non-zero exit code or was stopped by the system).
// case: When statefulset or kubevirt is restarted, it may enter the failed state for a short time,
// causing scall all to incorrectly reclaim the IP address, thereby changing the fixed IP address of the static Pod.

// PodPending means the pod has been accepted by the system, but one or more of the containers
// has not been started. This includes time before being bound to a node, as well as time spent
// pulling images onto the host.

// pod is running, pod has been assigned IP address

// The goal is to promptly reclaim IP addresses and to avoid having all trace data being blank when the spiderppol controller has just started or during a leader election.

// check pod status phase with its yaml

// handle same name pod with different uid in the ippool

// Check if the status.ips of the current K8S Pod has a value.
// If there is a value, it means that the pod has been started and the IP has been successfully assigned through cmdAdd
// If there is no value, it means that the new pod is still starting.

// handle the endpoint

// handle same name pod with different uid in the endpoint

// Check if the status.ips of the current K8S Pod has a value.
// If there is a value, it means that the pod has been started and the IP has been successfully assigned through cmdAdd
// If there is no value, it means that the new pod is still starting.

// Clean up outdated endpoints where both pod and owner reference no longer exist

func (s *SpiderGC) cleanOutdateEndpoint(ctx context.Context, suspiciousEndpointMap map[string]*spiderpoolv2beta1.WorkloadEndpointStatus) {
	_ = "STUB: not implemented"
	return
}

func (s *SpiderGC) checkEndpointExistInIPPool(ctx context.Context, epNamespace, epName, poolName string, logger *zap.Logger, status *spiderpoolv2beta1.WorkloadEndpointStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// Helps check if it is a valid static Pod (StatefulSet or Kubevirt), if it is a valid static Pod. Return true
func (s *SpiderGC) isValidStatefulsetOrKubevirt(ctx context.Context, logger *zap.Logger, podNS, podName, poolIP, ownerControllerType string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *SpiderGC) isShouldGCOrTraceStatelessTerminatingPodOnNode(ctx context.Context, pod *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	// check terminating Pod corresponding Node status
	return false, nil
}

// disable for gc terminating pod with Node Ready

// disable for gc terminating pod with Node NotReady

// shouldTraceOrReclaimIPInDeletionTimeStampPod check the deletion timestamp of the pod
// If the deletion timestamp of the pod is over, try to reclaim the IP
// If the deletion timestamp of the pod is not over and the pod still holds an IP, try to track the IP
// or the pod has no IP, try to reclaim the IP
func (s *SpiderGC) shouldTraceOrReclaimIPInDeletionTimeStampPod(scanAllLogger *zap.Logger, pod *corev1.Pod, shouldGcOrTraceStatelessTerminatingPod bool) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// The graceful deletion period of kubernetes Pod has not yet ended, and the Pod's already has an IP address. Let trace_worker track and recycle the IP in time.
// In addition, avoid that all trace data is blank when the controller is just started.
