// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationcontroller

import (
	"context"
	"time"

	"go.uber.org/zap"
	k8types "k8s.io/apimachinery/pkg/types"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	batchlisters "k8s.io/client-go/listers/batch/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/spiderpool/pkg/applicationcontroller/applicationinformers"
	"github.com/spidernet-io/spiderpool/pkg/election"
	"github.com/spidernet-io/spiderpool/pkg/subnetmanager"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

var logger *zap.Logger

type SubnetAppController struct {
	client    client.Client
	apiReader client.Reader

	subnetMgr     subnetmanager.SubnetManager
	workQueue     workqueue.RateLimitingInterface
	appController *applicationinformers.Controller

	deploymentsLister  appslisters.DeploymentLister
	deploymentInformer cache.SharedIndexInformer

	replicaSetLister   appslisters.ReplicaSetLister
	replicaSetInformer cache.SharedIndexInformer

	statefulSetLister   appslisters.StatefulSetLister
	statefulSetInformer cache.SharedIndexInformer

	daemonSetLister   appslisters.DaemonSetLister
	daemonSetInformer cache.SharedIndexInformer

	jobLister   batchlisters.JobLister
	jobInformer cache.SharedIndexInformer

	cronJobLister   batchlisters.CronJobLister
	cronJobInformer cache.SharedIndexInformer

	SubnetAppControllerConfig
}

type SubnetAppControllerConfig struct {
	EnableIPv4                    bool
	EnableIPv6                    bool
	AppControllerWorkers          int
	MaxWorkqueueLength            int
	WorkQueueMaxRetries           int
	WorkQueueRequeueDelayDuration time.Duration
	LeaderRetryElectGap           time.Duration
}

func NewSubnetAppController(client client.Client, apiReader client.Reader, subnetMgr subnetmanager.SubnetManager, subnetAppControllerConfig SubnetAppControllerConfig) (*SubnetAppController, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sac *SubnetAppController) SetupInformer(ctx context.Context, client kubernetes.Interface, leader election.SpiderLeaseElector) error {
	_ = "STUB: not implemented"
	return nil
}

func (sac *SubnetAppController) addEventHandlers(factory kubeinformers.SharedInformerFactory) error {
	_ = "STUB: not implemented"
	return nil
}

// Once we lost the leader but get leader later, we have to use a new workqueue.
// Because the former workqueue was already shut down and wouldn't be re-start forever.

// controllerAddOrUpdateHandler serves for kubernetes original controller applications(such as: Deployment,ReplicaSet,Job...),
// to create a new IPPool or scale the IPPool
func (sac *SubnetAppController) controllerAddOrUpdateHandler() applicationinformers.AppInformersAddOrUpdateFunc {
	_ = "STUB: not implemented"
	return *new(applicationinformers.AppInformersAddOrUpdateFunc)
}

// no need reconcile for HostNetwork application

// default IPAM mode

// no need reconcile for HostNetwork application

// check the app whether is the top controller or not

// default IPAM mode

// no need reconcile for HostNetwork application

// default IPAM mode

// no need reconcile for HostNetwork application

// check the app whether is the top controller or not

// default IPAM mode

// no need reconcile for HostNetwork application

// default IPAM mode

// no need reconcile for HostNetwork application

// default IPAM mode

// check the difference between the two object and choose to reconcile or not

// appWorkQueueKey involves application object meta namespaceKey and application kind
type appWorkQueueKey struct {
	MetaNamespaceKey string
	AppKind          string
	AppUID           k8types.UID
}

// enqueueApp will insert application custom appWorkQueueKey to the workQueue
func (sac *SubnetAppController) enqueueApp(ctx context.Context, obj interface{}, appKind string, appUID k8types.UID) {
	_ = "STUB: not implemented"
	return
}

// object meta key: 'namespace/name'

// validate workqueue capacity

func (sac *SubnetAppController) Run(stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (sac *SubnetAppController) runWorker() { _ = "STUB: not implemented"; return }

func (sac *SubnetAppController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// discard wrong input items

// requeue the conflict items

// if we set nonnegative number for the requeue delay duration, we will requeue it. otherwise we will discard it.

// syncHandler retrieves appWorkQueueKey from workQueue and try to create the auto-created IPPool or mark the IPPool status.AutoDesiredIPCount
func (sac *SubnetAppController) syncHandler(appKey appWorkQueueKey, log *zap.Logger) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// deployment.APIVersion is empty string

// replicaSet.APIVersion is empty string

// daemonSet.APIVersion is empty string

// statefulSet.APIVersion is empty string

// job.APIVersion is empty string

// cronJob.APIVersion is empty string

// applyAutoIPPool try to create an IPPool or mark IPPool desired IP number with the give SpiderSubnet configuration
func (sac *SubnetAppController) applyAutoIPPool(ctx context.Context, podSubnetConfig types.PodSubnetAnnoConfig,
	podController types.PodTopController, appReplicas int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// retrieve application pools

// We need to ignore the previous same NamespacedName application corresponding auto-created IPPool.
// Because this auto-created IPPool will be deleted by the system with 'ippool-reclaim'

// If 'ippool-reclaim' is true, we'll go to scale the auto-created IPPool because of the same application UID.
// If 'ippool-reclaim' is false, we'll reuse this auto-crated IPPool and refresh its application UID label.

// NewAggregate will check each the given error slice elements whether is nil or not

// hasSubnetConfigChanged checks whether application subnet configuration changed and the application replicas changed or not.
// The second parameter newSubnetConfig must not be nil.
func hasSubnetConfigChanged(ctx context.Context, oldSubnetConfig, newSubnetConfig *types.PodSubnetAnnoConfig,
	oldAppReplicas, newAppReplicas int,
) bool {
	_ = "STUB: not implemented"
	// go to reconcile directly with new application
	return false
}

// controllerDeleteHandler will return a function that clean up the application SpiderSubnet legacies (such as: the before created IPPools)
func (sac *SubnetAppController) controllerDeleteHandler() applicationinformers.APPInformersDelFunc {
	_ = "STUB: not implemented"
	return *new(applicationinformers.APPInformersDelFunc)
}

func (sac *SubnetAppController) deleteAutoPools(ctx context.Context, appUID k8types.UID) error {
	_ = "STUB: not implemented"
	return nil
}
