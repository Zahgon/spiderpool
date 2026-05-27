// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	"context"
	"time"

	v1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	frame "github.com/spidernet-io/e2eframework/framework"
	"github.com/spidernet-io/spiderpool/pkg/types"
	corev1 "k8s.io/api/core/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateIppool(f *frame.Framework, ippool *v1.SpiderIPPool, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// try to wait for finish last deleting

func BatchCreateIppoolWithSpecifiedIPNumber(frame *frame.Framework, ippoolNumber, ipNum int, isV4orv6Pool bool) (ipPoolNameList []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cycle create ippool

// traversal of ips in ip segment

// check whether the ip exists in ipMap

// If there is duplication in the middle, delete the dirty data

// continue back to OUTER_FOR to continue creating ippool

func DeleteIPPoolByName(f *frame.Framework, poolName string, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func GetIppoolByName(f *frame.Framework, poolName string) (*v1.SpiderIPPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetAllIppool(f *frame.Framework, opts ...client.ListOption) (*v1.SpiderIPPoolList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CheckIppoolForUsedIP(f *frame.Framework, ippool *v1.SpiderIPPool, podName, podNamespace string, ipAddrress *corev1.PodIP) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GetPodIPv4Address(pod *corev1.Pod) *corev1.PodIP { _ = "STUB: not implemented"; return nil }

func GetPodIPv6Address(pod *corev1.Pod) *corev1.PodIP { _ = "STUB: not implemented"; return nil }

func CheckPodIPRecordInIPPool(f *frame.Framework, v4IppoolNameList, v6IppoolNameList []string, podList *corev1.PodList) (allIPRecorded, noneIPRecorded, partialIPRecorded bool, err error) {
	_ = "STUB: not implemented"
	return false, false, false, nil
}

func GetNamespaceDefaultIppool(f *frame.Framework, namespace string) (v4IppoolList, v6IppoolList []string, e error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GetWorkloadByName(f *frame.Framework, namespace, name string) (*v1.SpiderEndpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CheckIppoolForPodName(f *frame.Framework, ippool *v1.SpiderIPPool, podName, podNamespace string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func WaitIPReclaimedFinish(f *frame.Framework, v4IppoolNameList, v6IppoolNameList []string, podList *corev1.PodList, timeOut time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func GenerateExampleIpv4poolObject(ipNum int) (string, *v1.SpiderIPPool) {
	_ = "STUB: not implemented"
	return "", nil
}

func PatchConfigMap(f *frame.Framework, oldcm, newcm *corev1.ConfigMap, opts ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func GenerateExampleIpv6poolObject(ipNum int) (string, *v1.SpiderIPPool) {
	_ = "STUB: not implemented"
	return "", nil
}

func DeleteIPPoolUntilFinish(f *frame.Framework, poolName string, ctx context.Context, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateIppool(f *frame.Framework, ippool *v1.SpiderIPPool, opts ...client.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func PatchIppool(f *frame.Framework, desiredPool, originalPool *v1.SpiderIPPool, opts ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func BatchDeletePoolUntilFinish(f *frame.Framework, iPPoolNameList []string, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func GenerateRandomIPV4() string { _ = "STUB: not implemented"; return "" }

func GenerateRandomIPV6() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Waiting for Ippool Status Condition By Allocated IPs meets expectations
// can be used to detect dirty IPs recorded in ippool to be reclaimed automatically
func WaitIppoolStatusConditionByAllocatedIPs(ctx context.Context, f *frame.Framework, poolName, checkIPs string, isRecord bool) error {
	_ = "STUB: not implemented"
	return nil
}

// When the Pod IP resource is reclaimed, wait for the corresponding workload deletion to complete
func WaitWorkloadDeleteUntilFinish(ctx context.Context, f *frame.Framework, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckUniqueUUIDInSpiderPool(f *frame.Framework, poolName string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetIppoolsInSubnet(f *frame.Framework, subnetName string) (*v1.SpiderIPPoolList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetPoolNameListInSubnet(f *frame.Framework, subnetName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateIppoolInSpiderSubnet(ctx context.Context, f *frame.Framework, subnetName string, pool *v1.SpiderIPPool, ipNum int) error {
	_ = "STUB: not implemented"
	return nil
}

// The informer of SpiderSubnet will delay synchronizing its own state information,
// and build SpiderIPPool concurrently to add a retry mechanism to handle dirty reads.

// BatchCreateIPPoolsInSpiderSubnet will create a set of identical versions of ippools for you under the desired subnet.
// subnet and subnetIPRanges must belong to the same IP version.
func BatchCreateIPPoolsInSpiderSubnet(f *frame.Framework, version types.IPVersion, subnet string, subnetIPRanges []string, poolNum, ipNum int) (poolNameList []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
GetPodIPAddressFromIppool is to get the IP from the ippool by name and namespace.
when the application has multiple IPs, but only one is displayed in the application's status,
it is necessary to get the one from the ippool to compare with the actual one.
*/
func GetPodIPAddressFromIppool(f *frame.Framework, poolName, namespace, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func WaitWebhookReady(ctx context.Context, f *frame.Framework, webhookPort string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetSpiderControllerEnvValue(f *frame.Framework, envName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type NetworkStatus struct {
	Name      string   `json:"name"`
	Interface string   `json:"interface"`
	IPs       []string `json:"ips"`
	MAC       string   `json:"mac"`
	Default   bool     `json:"default"`
}

// ParsePodNetworkAnnotation parses the 'PodMultusNetworksStatus' annotation from the given Pod
// and extracts all IP addresses associated with the Pod's network interfaces.
func ParsePodNetworkAnnotation(f *frame.Framework, pod *corev1.Pod) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal the JSON from the annotation into a slice of NetworkStatus

// CheckIppoolSanity checks the integrity and correctness of the IP pool's allocation status.
// It ensures that each IP in the pool is:
// 1. Allocated to a single Pod, whose UID matches the record in the IP pool.
// 2. Correctly tracked by its associated endpoint, confirming that the Pod's UID matches the endpoint's UID.
// 3. The actual number of IPs in use is compared against the IP pool's reported usage count to ensure consistency.
func CheckIppoolSanity(f *frame.Framework, poolName string) error {
	_ = "STUB: not implemented"
	// Retrieve the IPPool by name
	return nil
}

// Parse the allocated IPs from the IP pool status

// Track the actual number of IPs in use

// The total number of assigned IP addresses

// Split the pod NamespacedName to get the namespace and pod name

// Retrieve the Pod object by its name and namespace

// The status of IPPool is automatically synchronized by the IPPool informer based on the events it receives.
// In the CI environment, the creation of IPPools happens very quickly, and their health checks are performed promptly.
// When checking the TotalIPCount status, if the spiderpool-controller undergoes a leader election or the informer has not yet completed synchronization,
// the IPPool status TotalIPCount may be nil. This can lead to a panic.
// In such cases, try waiting for the informer to complete status synchronization before checking the robustness of the IPPool.

// Ensure that the IP pool's reported usage matches the actual usage
