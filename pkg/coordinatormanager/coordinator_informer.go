// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package coordinatormanager

import (
	"context"
	"time"

	ciliumLister "github.com/cilium/cilium/pkg/k8s/client/listers/cilium.io/v2alpha1"
	"go.uber.org/zap"
	coreinformers "k8s.io/client-go/informers/core/v1"
	networkingInformer "k8s.io/client-go/informers/networking/v1alpha1"
	"k8s.io/client-go/kubernetes"
	corelister "k8s.io/client-go/listers/core/v1"
	networkingLister "k8s.io/client-go/listers/networking/v1alpha1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/spiderpool/pkg/election"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	clientset "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	spiderinformers "github.com/spidernet-io/spiderpool/pkg/k8s/client/informers/externalversions/spiderpool.spidernet.io/v2beta1"
	spiderlisters "github.com/spidernet-io/spiderpool/pkg/k8s/client/listers/spiderpool.spidernet.io/v2beta1"
)

const (
	auto    = "auto"
	cluster = "cluster"
	calico  = "calico"
	cilium  = "cilium"
	flannel = "flannel"
	none    = "none"
)

var SupportedPodCIDRType = []string{auto, cluster, calico, cilium, none}

const (
	calicoIPPoolCRDName = "ippools.crd.projectcalico.org"
	ciliumIPPoolCRDName = "ciliumpodippools.cilium.io"
	ciliumConfig        = "cilium-config"
	kubeadmConfigMap    = "kubeadm-config"
)

const (
	NotReady = "NotReady"
	Synced   = "Synced"
)

const messageEnqueueCoordiantor = "Enqueue Coordinator"

var InformerLogger *zap.Logger

type CoordinatorController struct {
	K8sClient *kubernetes.Clientset
	Manager   ctrl.Manager
	Client    client.Client
	APIReader client.Reader

	CoordinatorLister spiderlisters.SpiderCoordinatorLister
	ConfigmapLister   corelister.ConfigMapLister
	// only not to nil if the k8s serviceCIDR is enabled
	ServiceCIDRLister networkingLister.ServiceCIDRLister
	// only not to nil if the cilium multu-pool is enabled
	CiliumIPPoolLister ciliumLister.CiliumPodIPPoolLister

	CoordinatorSynced cache.InformerSynced
	ConfigmapSynced   cache.InformerSynced
	// only not nil if the k8s serviceCIDR is enabled
	ServiceCIDRSynced cache.InformerSynced
	// only not to nil if the cilium multu-pool is enabled
	CiliumIPPoolsSynced cache.InformerSynced

	Workqueue workqueue.RateLimitingInterface

	LeaderRetryElectGap time.Duration
	ResyncPeriod        time.Duration

	DefaultCniConfDir      string
	CiliumConfigMap        string
	DefaultCoordinatorName string
}

func (cc *CoordinatorController) SetupInformer(
	ctx context.Context,
	spiderClientset clientset.Interface,
	k8sClientset *kubernetes.Clientset,
	leader election.SpiderLeaseElector,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) StartWatchPodCIDR(ctx context.Context, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) addEventHandlers(
	coordinatorInformer spiderinformers.SpiderCoordinatorInformer,
	configmapInformer coreinformers.ConfigMapInformer,
	serviceCIDRInformer networkingInformer.ServiceCIDRInformer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) addServiceCIDRHandler(serviceCIDRInformer cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) enqueueCoordinatorOnAdd(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cc *CoordinatorController) enqueueCoordinatorOnUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cc *CoordinatorController) enqueueCoordinatorOnConfigMapAdd(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cc *CoordinatorController) enqueueCoordinatorOnConfigMapUpdated(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cc *CoordinatorController) enqueueCoordinatorOnConfigMapDeleted(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (cc *CoordinatorController) run(ctx context.Context, workers int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) runWorker(ctx context.Context) { _ = "STUB: not implemented"; return }

func (cc *CoordinatorController) processNextWorkItem(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (cc *CoordinatorController) syncHandler(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) updatePodAndServerCIDR(ctx context.Context, logger *zap.Logger, coord *spiderpoolv2beta1.SpiderCoordinator) *spiderpoolv2beta1.SpiderCoordinator {
	_ = "STUB: not implemented"
	return nil
}

// TODO(@cyclinder): Do we need watch if /etc/cni/net.d has changed?

// try to get ClusterCIDR from kubeadm-config ConfigMap

// Success to get ClusterCIDR from kubeadm-config

// if kubeadm-config ConfigMap not found, try to get ClusterCIDR from kube-controller-manager Pod

func (cc *CoordinatorController) updateCalicoPodCIDR(ctx context.Context, coordinator *spiderpoolv2beta1.SpiderCoordinator) error {
	_ = "STUB: not implemented"
	return nil
}

// sort the list to admit podCIDR has changed.

func (cc *CoordinatorController) WatchCalicoIPPools(ctx context.Context, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) WatchCiliumIPPools(ctx context.Context, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) updateCiliumPodCIDR(k8sPodCIDR []string, coordinator *spiderpoolv2beta1.SpiderCoordinator) error {
	_ = "STUB: not implemented"
	return nil
}

func (cc *CoordinatorController) fetchCiliumIPPools(coordinator *spiderpoolv2beta1.SpiderCoordinator) error {
	_ = "STUB: not implemented"
	return nil
}

// sort the list to admit podCIDR has changed.

func (cc *CoordinatorController) updateServiceCIDR(logger *zap.Logger, coordCopy *spiderpoolv2beta1.SpiderCoordinator) error {
	_ = "STUB: not implemented"
	// fetch kubernetes ServiceCIDR
	return nil
}

// serviceCIDR feature is disable if ServiceCIDRLister is nil

// sort the list to admit serviceCIDR has changed due to the order.

func fetchType(cniDir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func setStatus2NoReady(logger *zap.Logger, reason string, copy *spiderpoolv2beta1.SpiderCoordinator) {
	_ = "STUB: not implemented"
	return
}
