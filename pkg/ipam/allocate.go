// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ipam

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"github.com/spidernet-io/spiderpool/api/v1/agent/models"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

func (i *ipam) Allocate(ctx context.Context, addArgs *models.IpamAddArgs) (*models.IpamAddResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Flag to indicate whether outdated IPs should be released.
// Check if StatefulSets are enabled in the configuration and
// if the pod's top controller is a StatefulSet.
// If an endpoint exists, attempt to release outdated IPs for
// the StatefulSet if necessary, and return an error if the
// operation fails.

// If the endpoint previously existed and its UID is not the same as the Pod's UID,
// the outdated endpoint may have already been deleted. In this case, we must set
// it to nil. Otherwise, we might not create a new endpoint in the PatchIPAllocationResults
// function, which could result in the Pod's endpoint object being lost.

func (i *ipam) releaseStsOutdatedIPIfNeed(ctx context.Context, addArgs *models.IpamAddArgs,
	pod *corev1.Pod, endpoint *spiderpoolv2beta1.SpiderEndpoint, podTopController types.PodTopController, isMultipleNicWithNoName bool,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Spiderpool assigns IP addresses to NICs one by one.
// Some NICs may have their IP pools changed, while others may remain unchanged.
// Record these changes and differences to handle specific NICs accordingly.

// If using the multi-NIC feature through ipam.spidernet.io/ippools without specifying interface names,
// and if the IP pool of one NIC changes, only reclaiming the corresponding endpoint IP could cause the IPAM allocation method to lose allocation records.
// When the interface name is not specified, the allocated NIC name might be "", which cannot be handled properly.
// If isMultipleNicWithNoName is true, all NIC IP addresses will be reclaimed and reallocated.

// All other cases determine here whether an IP address needs to be reclaimed.

// The multi-NIC feature can be used in the following two ways:
//   1. By specifying additional NICs through k8s.v1.cni.cncf.io/networks and configure the default pool.
//   2. By using ipam.spidernet.io/ippools (excluding cases where the interface name is empty).
// When a change is detected in the corresponding NIC's IP pool,
// the IP information for that NIC will be automatically reclaimed and reallocated.

// According to the NIC allocation mechanism, we check whether the pool information for each NIC has changed.
// If there is no change, we do not need to reclaim the corresponding endpoint and IP for that NIC.

// The endpoint should be deleted in the following cases:
// 1. If the multi-NIC feature is used through ipam.spidernet.io/ippools without specifying the interface name, and the IPPool has changed.
// 2. In other multi-NIC or single-NIC scenarios, if the IPPool of all NICs has changed.

// Only update the endpoint and IP corresponding to the changed NIC.

func (i *ipam) retrieveStaticIPAllocation(ctx context.Context, nic string, pod *corev1.Pod, endpoint *spiderpoolv2beta1.SpiderEndpoint) (*models.IpamAddResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The first allocation or multi-NIC.

func (i *ipam) reallocateIPPoolIPRecords(ctx context.Context, uid string, endpoint *spiderpoolv2beta1.SpiderEndpoint) error {
	_ = "STUB: not implemented"
	return nil
}

// Record the metric of queuing time for allocating.

func (i *ipam) retrieveExistingIPAllocation(ctx context.Context, uid, nic string, endpoint *spiderpoolv2beta1.SpiderEndpoint, isMultipleNicWithNoName bool) (*models.IpamAddResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create -> Delete -> Create a Pod with the same namespace and name in
// a short time will cause some unexpected phenomena discussed in
// https://github.com/spidernet-io/spiderpool/issues/1187.

// In AI training scenarios using Jobs, frequent creation failures may
// cause new Pods to fail to start. For cases where the endpoint UUID
// and Pod UID are inconsistent, Maybe be better to retain the
// endpoint object and patch the status later in PatchIPAllocationResults
// function, instead of deleting the endpoint object.
// see https://github.com/spidernet-io/spiderpool/issues/4916

// update Endpoint NIC name in multiple NIC with no name mode by annotation "ipam.spidernet.io/ippools"

func (i *ipam) allocateInStandardMode(ctx context.Context, addArgs *models.IpamAddArgs, pod *corev1.Pod, endpoint *spiderpoolv2beta1.SpiderEndpoint, podController types.PodTopController) (*models.IpamAddResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort the results in order by NIC sequence in multiple NIC with no name specified mode

// replace the first NIC name from "0" to "eth0"

// replace the routes NIC name from "0" to "eth0"

// Actually in allocate Standard Mode, we just need the current turn NIC allocation result,
// but here are the all NICs results

func (i *ipam) genToBeAllocatedSet(ctx context.Context, addArgs *models.IpamAddArgs, pod *corev1.Pod, podController types.PodTopController) (ToBeAllocateds, error) {
	_ = "STUB: not implemented"
	return *new(ToBeAllocateds), nil
}

// sort IPPool candidates

func (i *ipam) allocateIPsFromAllCandidates(ctx context.Context, tt ToBeAllocateds, pod *corev1.Pod, podController types.PodTopController) ([]*types.AllocationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Record the metric of queuing time for allocating.

// the results are not in order by the NIC sequence right now

func (i *ipam) allocateIPFromCandidate(ctx context.Context, c *PoolCandidate, nic string, cleanGateway bool, pod *corev1.Pod, podController types.PodTopController) (*types.AllocationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *ipam) precheckPoolCandidates(ctx context.Context, t *ToBeAllocated) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *ipam) filterPoolCandidates(ctx context.Context, t *ToBeAllocated, pod *corev1.Pod, podTopController types.PodTopController, addArgs *models.IpamAddArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *ipam) selectByPod(ctx context.Context, version types.IPVersion, ipPool *spiderpoolv2beta1.SpiderIPPool, pod *corev1.Pod, podTopController types.PodTopController, nic string, podNetNamespace *string, matchMasterSubnet bool) error {
	_ = "STUB: not implemented"
	return nil
}

// node

// namespace

// pod affinity

// the first NIC

// default net-attach-def specified in the annotations

// Use ParsePodNetworkAnnotation (not ParsePodNetworkObjectName) to handle
// both JSON and comma-delimited formats, consistent with upstream multus-cni.

// Reference from Multus source codes: The CRD object of default network should only be defined in multusNamespace
// In multus, multusNamespace serves for (clusterNetwork/defaultNetworks)

// the additional NICs must own a Multus CR object

// We regard the NIC name was specified by the user for the previous judgement.
// For the latter judgement(multiple NIC with no name specified mode), we just need to check whether the sequence is same with the net-attach-def resource

// Refer from the multus-cni source codes, for annotation "k8s.v1.cni.cncf.io/networks" value without Namespace,
// we will regard the pod Namespace as the value's namespace

// impossible

// multus

// for the ippool.spec.multusName property, if the user doesn't specify the net-attach-def resource namespace,
// we'll regard it in the Spiderpool installation namespace

// filter out ippools by checking the subnet is matched with the pod's node network zone
// only for sriov cni

func (i *ipam) filterPoolCandidatesByPfSubnet(pool *spiderpoolv2beta1.SpiderIPPool, podNetNamespace, nic string) error {
	_ = "STUB: not implemented"
	return nil
}

// impossible

func (i *ipam) verifyPoolCandidates(tt ToBeAllocateds) error {
	_ = "STUB: not implemented"
	//	for _, t := range tt {
	//		var allIPPools []*spiderpoolv2beta1.SpiderIPPool
	//		for _, c := range t.PoolCandidates {
	//			allIPPools = append(allIPPools, c.PToIPPool.IPPools()...)
	//		}
	//
	//		vlanToPools := map[types.Vlan][]string{}
	//		for _, ipPool := range allIPPools {
	//			vlanToPools[*ipPool.Spec.Vlan] = append(vlanToPools[*ipPool.Spec.Vlan], ipPool.Name)
	//		}
	//
	//		if len(vlanToPools) > 1 {
	//			return fmt.Errorf("%w, the VLANs of the IPPools corresponding to NIC %s are not all the same: %v", constant.ErrWrongInput, t.NIC, vlanToPools)
	//		}
	//	}
	return nil
}

// IsDetectGatewayReachableForKubeVirtPod disable IP conflict detection for the kubevirt vm live migration pod,
// If we don't do this, it will cause the migration pod never be started.
func (i *ipam) IsDetectGatewayReachableForKubeVirtPod(ctx context.Context, pod *corev1.Pod) (enableIPConflictDetection bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// disable IP conflict detection for the kubevirt vm live migration pod
// return directly if not a kubevirt vm pod

// the live migration new pod has the annotation "kubevirt.io/migrationJobName"
// we just only cancel IP conflict detection for the live migration new pod.

// kubevirt vm pod corresponding SpiderEndpoint uses kubevirt VM/VMI name

// cancel IP conflict detection because there's a moment the old vm pod still running during the vm live migration phase

// if we don't found the kubevirt migrated vm pod, still execute IP conflict detection

// sortPoolCandidates would sort IPPool candidates sequence depends on the IPPool multiple affinities.
func sortPoolCandidates(preliminary ToBeAllocateds) { _ = "STUB: not implemented"; return }

// new IPPool candidate names

// collect all IPPool resource from PoolCandidate.PToIPPool

// make it order with ippoolmanager.ByPoolPriority interface rules

// set the new IPPool candidate names to PoolCandidate.Pools
