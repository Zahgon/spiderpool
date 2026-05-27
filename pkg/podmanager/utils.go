// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package podmanager

import (
	"context"

	v2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	crdclientset "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	"github.com/spidernet-io/spiderpool/pkg/namespacemanager"
	corev1 "k8s.io/api/core/v1"
)

func IsPodAlive(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// IsStaticIPPod checks the given pod's controller ownerReference whether is StatefulSet or KubevirtVMI
func IsStaticIPPod(enableStatefulSet, enableKubevirtStaticIP bool, pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

// podNetworkMutatingWebhook handles the mutating webhook for pod networks.
// It checks if the pod has the required label for mutation, retrieves the corresponding
// SpiderMultusConfigs, and injects the network configuration into the pod.
//
// Parameters:
//   - apiReader: A client.Reader interface for accessing Kubernetes API objects
//   - pod: A pointer to the corev1.Pod object to be mutated
//
// Returns:
//   - An error if any step in the process fails, nil otherwise
func podNetworkMutatingWebhook(ctx context.Context, spiderClient crdclientset.Interface, nsManager namespacemanager.NamespaceManager, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func needPodNetworkInjection(ctx context.Context, nsManager namespacemanager.NamespaceManager, pod *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getEffectiveResourceInjectValue(ctx context.Context, nsManager namespacemanager.NamespaceManager, pod *corev1.Pod, anno string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func getMultusConfigSortKey(mc v2beta1.SpiderMultusConfig) string {
	_ = "STUB: not implemented"
	return ""
}

func sortMultusConfigs(multusConfigs *v2beta1.SpiderMultusConfigList) {
	_ = "STUB: not implemented"
	return
}

// InjectPodNetwork injects network configurations into the pod based on the provided SpiderMultusConfigs.
// It checks for CNI type consistency, updates the pod's network attachment annotations,
// and prepares a map of resources to be injected.
//
// Parameters:
//   - pod: A pointer to the corev1.Pod object to be updated
//   - multusConfigs: A list of SpiderMultusConfig objects to be applied to the pod
//
// Returns:
//   - An error if there's an inconsistency in CNI types, nil otherwise
func InjectPodNetwork(pod *corev1.Pod, multusConfigs v2beta1.SpiderMultusConfigList) error {
	_ = "STUB: not implemented"
	return nil
}

// InjectRdmaResourceToPod injects RDMA resources into the pod's containers.
// It checks each container for existing resource requests/limits and updates
// the resourceMap accordingly. If a resource is not found in any container,
// it is injected into the first container's resource requests.
//
// Parameters:
//   - resourceMap: A map of resource names to boolean values indicating if they've been found
//   - pod: A pointer to the corev1.Pod object to be updated
func InjectRdmaResourceToPod(resourceMap map[string]bool, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// the resource has found in pod, skip

// try to find the resource in container resources.limits

func DoValidateRdmaResouce(mc v2beta1.SpiderMultusConfig) error {
	_ = "STUB: not implemented"
	return nil
}
