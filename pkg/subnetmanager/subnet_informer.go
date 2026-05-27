// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package subnetmanager

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/spiderpool/pkg/election"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	clientset "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	informers "github.com/spidernet-io/spiderpool/pkg/k8s/client/informers/externalversions/spiderpool.spidernet.io/v2beta1"
	listers "github.com/spidernet-io/spiderpool/pkg/k8s/client/listers/spiderpool.spidernet.io/v2beta1"
	spiderpooltypes "github.com/spidernet-io/spiderpool/pkg/types"
)

const (
	messageEnqueueSubnet = "Enqueue Subnet"
	messageWorkqueueFull = "Workqueue is full, dropping the element"
)

var InformerLogger *zap.Logger

type SubnetController struct {
	Client    client.Client
	APIReader client.Reader

	SubnetsLister listers.SpiderSubnetLister
	IPPoolsLister listers.SpiderIPPoolLister

	SubnetIndexer cache.Indexer
	IPPoolIndexer cache.Indexer

	SubnetsSynced cache.InformerSynced
	IPPoolsSynced cache.InformerSynced

	Workqueue workqueue.RateLimitingInterface

	LeaderRetryElectGap     time.Duration
	ResyncPeriod            time.Duration
	SubnetControllerWorkers int
	MaxWorkqueueLength      int

	DynamicClient    dynamic.Interface
	dynamicFactory   dynamicinformer.DynamicSharedInformerFactory
	dynamicWorkqueue workqueue.RateLimitingInterface
	recordedResource sync.Map
}

type thirdControllerKey struct {
	MetaNamespaceKey string
	AppUID           types.UID
}

func (sc *SubnetController) SetupInformer(ctx context.Context, client clientset.Interface, leader election.SpiderLeaseElector) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SubnetController) addEventHandlers(subnetInformer informers.SpiderSubnetInformer, ipPoolInformer informers.SpiderIPPoolInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SubnetController) enqueueSubnetOnAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (sc *SubnetController) enqueueSubnetOnUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// enqueueSubnetOnIPPoolChange receives the IPPool resources events
func (sc *SubnetController) enqueueSubnetOnIPPoolChange(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (sc *SubnetController) run(ctx context.Context, workers int) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SubnetController) runWorker(ctx context.Context) { _ = "STUB: not implemented"; return }

func (sc *SubnetController) processNextWorkItem(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (sc *SubnetController) runDynamicWorker(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (sc *SubnetController) processDynamicNextWorkItem(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (sc *SubnetController) syncHandler(ctx context.Context, subnetName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// syncMetadata add "ipam.spidernet.io/subnet-cidr" label for the SpiderSubnet object
func (sc *SubnetController) syncMetadata(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) error {
	_ = "STUB: not implemented"
	return nil
}

// syncControllerSubnet would set ownerReference and add "ipam.spidernet.io/owner-spider-subnet" label for the previous orphan IPPool
func (sc *SubnetController) syncControllerSubnet(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) error {
	_ = "STUB: not implemented"
	return nil
}

func (sc *SubnetController) syncControlledIPPoolIPs(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) error {
	_ = "STUB: not implemented"
	return nil
}

// Merge pre-allocated IP addresses of each IPPool and calculate their count.

// Only auto-created IPPools have the field 'Application'.

// discard this invalid allocation for subnet.status

// if it's a third-party controller, we'll watch its deletion hook function to GC SpiderSubnet.status

// discard the legacy allocation for subnet.status

// discard this invalid allocation

// record the metric of how many IPPools the Subnet has.

// Update the count of total IP addresses.

func (sc *SubnetController) removeFinalizer(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) error {
	_ = "STUB: not implemented"
	return nil
}

// Some IP addresses are still occupied by the controlled IPPools, ignore
// to remove the finalizer.

func (sc *SubnetController) monitorThirdController(ctx context.Context, appNamespacedName spiderpooltypes.AppNamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

// if stopped, let's clean up the cache directly
