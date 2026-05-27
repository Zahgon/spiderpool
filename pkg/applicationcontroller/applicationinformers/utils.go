// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationinformers

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apitypes "k8s.io/apimachinery/pkg/types"
	k8svalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/client"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

// ClusterSubnetAutoPoolDefaultRedundantIPNumber is a singleton recording cluster subnet AutoPool default redundant IP number.
var ClusterSubnetAutoPoolDefaultRedundantIPNumber = new(int)

var errInvalidInput = func(str string) error {
	return fmt.Errorf("invalid input '%s'", str)
}

const (
	maxNameLength = k8svalidation.DNS1123SubdomainMaxLength
	randomLength  = 5
)

func AutoPoolName(controllerName string, ipVersion types.IPVersion, ifName string, appUID apitypes.UID) string {
	_ = "STUB: not implemented"
	// the format of uuid is "xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	// ref: https://github.com/google/uuid/blob/44b5fee7c49cf3bcdf723f106b36d56ef13ccc88/uuid.go#L185
	return ""
}

// we constrain the max random length to 5

// 7 means "auto${IPVersion}" prefix and 2 bound symbol "-" length

// ApplicationLabelGV switches the kubernetes APIVersion from "/" link format to "_" link format for kubernetes label value usage.
func ApplicationLabelGV(apiVersion string) string {
	_ = "STUB: not implemented"
	// Kubernetes API Group might be empty, ref: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#api-versions
	return ""
}

// ParseApplicationGVStr will parse a label value string back to apiVersion format string, its corresponding function is ApplicationLabelGV.
func ParseApplicationGVStr(str string) (apiVersion string, isMatch bool) {
	_ = "STUB: not implemented"
	return "", false

	// no API Group
}

// ApplicationNamespacedName will joint the application apiVersion, application type, namespace and name as a string, then we need unpack it for tracing
// [ns and object name constraint Ref]: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/
// We set format is "{apiVersion}:{appKind}:{appNS}:{appName}"
func ApplicationNamespacedName(appNamespacedName types.AppNamespacedName) string {
	_ = "STUB: not implemented"
	// Kubernetes API Group might be empty, ref: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#api-versions
	return ""
}

// ParseApplicationNamespacedName will unpack the appNamespacedNameKey, its corresponding function is ApplicationNamespacedName
func ParseApplicationNamespacedName(appNamespacedNameKey string) (appNamespacedName types.AppNamespacedName, isMatch bool) {
	_ = "STUB: not implemented"
	return *new(types.AppNamespacedName), false
}

// no API Group

func GetAppReplicas(replicas *int32) int { _ = "STUB: not implemented"; return 0 }

func GenSubnetFreeIPs(subnet *spiderpoolv2beta1.SpiderSubnet) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSubnetAnnoConfig generates SpiderSubnet configuration from pod annotation,
// if the pod doesn't have the related subnet annotation but has IPPools/IPPool relative annotation it will return nil.
// If the pod doesn't have any subnet/ippool annotations, it will use the cluster default subnet configuration.
func GetSubnetAnnoConfig(podAnnotations map[string]string, log *zap.Logger) (*types.PodSubnetAnnoConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// annotation: ipam.spidernet.io/subnets

// annotation: ipam.spidernet.io/subnet

// annotation: ipam.spidernet.io/ippool-ip-number

// check out negative number

// annotation: "ipam.spidernet.io/reclaim-ippool", reclaim IPPool or not (default true)

// mutateAndValidateSubnetAnno will filter multiple subnets you specified and only leaves you the first one to use.
// And it also checks Interface name or subnets you specified whether are duplicate.
func mutateAndValidateSubnetAnno(subnetConfig *types.PodSubnetAnnoConfig) error {
	_ = "STUB: not implemented"
	// the present version, we just only support one SpiderSubnet object to choose
	return nil
}

// all none

// validate duplicate subnet

// validate duplicate interface

// all none

// specify 'eth0' as the default single interface if it's none.

// GetPoolIPNumber judges the given parameter is fixed or flexible
func GetPoolIPNumber(str string) (isFlexible bool, ipNum int, err error) {
	_ = "STUB: not implemented"

	// the '+' sign counts must be '0' or '1'
	return false, 0, nil
}

// CalculateJobPodNum will calculate the job replicas
// once Parallelism and Completions are unset, the API-server will set them to 1
// reference: https://kubernetes.io/docs/concepts/workloads/controllers/job/
func CalculateJobPodNum(jobSpecParallelism, jobSpecCompletions *int32) int {
	_ = "STUB: not implemented"
	return 0
}

// parallel Jobs with a work queue

// ignore negative integer, cause API-server will refuse the job creation

// non-parallel Jobs

// ignore negative integer, cause API-server will refuse the job creation

// parallel Jobs with a fixed completion count

// ignore negative integer, cause API-server will refuse the job creation

// IsDefaultIPPoolMode judges whether we use subnet feature or not with the given parameter types.PodSubnetAnnoConfig
func IsDefaultIPPoolMode(subnetConfig *types.PodSubnetAnnoConfig) bool {
	_ = "STUB: not implemented"
	return false
}

// SpiderSubnet with multiple interfaces

// SpiderSubnet with single interface

// containsDuplicate checks whether the given string array has the duplicate element
func containsDuplicate(arr []string) bool { _ = "STUB: not implemented"; return false }

// ShouldReclaimIPPool will check pod annotation "ipam.spidernet.io/ippool-reclaim"
func ShouldReclaimIPPool(anno map[string]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// no specified reclaim-IPPool, default to set it true

// IsAppExist will check the application whether exists or not. If it exists, it will return the application corresponding UID
func IsAppExist(ctx context.Context, cacheClient client.Client, dynamicClient dynamic.Interface, appNamespacedName types.AppNamespacedName) (isExist bool, appUID apitypes.UID, err error) {
	_ = "STUB: not implemented"
	return false, *new(apitypes.UID), nil
}

// if the application is no longer exist, we should delete the IPPool

func GenerateGVR(appNamespacedName types.AppNamespacedName) (schema.GroupVersionResource, error) {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionResource), nil
}

func IsThirdController(appNamespacedName types.AppNamespacedName) bool {
	_ = "STUB: not implemented"
	return false
}

func IsReclaimAutoPoolLabelValue(isReclaim bool) string { _ = "STUB: not implemented"; return "" }

func AutoPoolIPVersionLabelValue(ipVersion types.IPVersion) string {
	_ = "STUB: not implemented"
	return ""
}

// HasSubnetsAnnotation checks the given annotation whether contains 'ipam.spidernet.io/subnets' or 'ipam.spidernet.io/subnet' key-value pair
func HasSubnetsAnnotation(anno map[string]string) bool { _ = "STUB: not implemented"; return false }
