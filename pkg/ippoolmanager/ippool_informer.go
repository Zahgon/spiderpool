// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package ippoolmanager

import (
	"context"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/spiderpool/pkg/election"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	crdclientset "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	informers "github.com/spidernet-io/spiderpool/pkg/k8s/client/informers/externalversions/spiderpool.spidernet.io/v2beta1"
	listers "github.com/spidernet-io/spiderpool/pkg/k8s/client/listers/spiderpool.spidernet.io/v2beta1"
)

var informerLogger *zap.Logger

type IPPoolController struct {
	IPPoolControllerConfig
	client        client.Client
	dynamicClient dynamic.Interface
	poolLister    listers.SpiderIPPoolLister
	poolSynced    cache.InformerSynced
	poolWorkqueue workqueue.RateLimitingInterface
}

type IPPoolControllerConfig struct {
	IPPoolControllerWorkers       int
	EnableSpiderSubnet            bool
	EnableAutoPoolForApplication  bool
	MaxWorkqueueLength            int
	WorkQueueMaxRetries           int
	LeaderRetryElectGap           time.Duration
	WorkQueueRequeueDelayDuration time.Duration
	ResyncPeriod                  time.Duration
}

func NewIPPoolController(poolControllerConfig IPPoolControllerConfig, client client.Client, dynamicClient dynamic.Interface) *IPPoolController {
	_ = "STUB: not implemented"
	return nil
}

func (ic *IPPoolController) SetupInformer(ctx context.Context, client crdclientset.Interface, controllerLeader election.SpiderLeaseElector) error {
	_ = "STUB: not implemented"
	return nil
}

func (ic *IPPoolController) addEventHandlers(poolInformer informers.SpiderIPPoolInformer) error {
	_ = "STUB: not implemented"
	return nil
}

// for all IPPool processing

// enqueueIPPool will check the given pool and enqueue them into different workqueue
func (ic *IPPoolController) enqueueIPPool(obj interface{}) { _ = "STUB: not implemented"; return }

// the Normal IPPools enqueue the corresponding NormalPoolWorkqueue

// Run will set up the event handlers for IPPool, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (ic *IPPoolController) Run(stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// worker runs a worker thread that just dequeues items, processes them, and marks them done.
// This will update SpiderIPPool status counts
func (ic *IPPoolController) runWorker() { _ = "STUB: not implemented"; return }

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it with the given function handler.
// the processNextWorkItem is never invoked concurrently with the same key.
func (ic *IPPoolController) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

// The IPPool resource may no longer exist, in which case we stop
// processing.

// discard some wrong input items

// if we set nonnegative number for the requeue delay duration, we will requeue it. otherwise we will discard it.

func (ic *IPPoolController) handleIPPool(ctx context.Context, pool *spiderpoolv2beta1.SpiderIPPool) (err error) {
	_ = "STUB: not implemented"
	// checkout the Auto-created IPPools whether need to scale or clean up legacies
	return nil
}

// update the IPPool status properties

// metrics

// syncHandler will calculate and update the provided SpiderIPPool status AllocatedIPCount or TotalIPCount.
// And it will also remove finalizer once the IPPool is dying and no longer being used.
func (ic *IPPoolController) syncHandler(ctx context.Context, pool *spiderpoolv2beta1.SpiderIPPool) error {
	_ = "STUB: not implemented"
	// remove finalizer to delete the dying IPPool when the IPPool is no longer being used
	return nil
}

// initial the original data

// removeFinalizer removes SpiderIPPool finalizer
func (ic *IPPoolController) removeFinalizer(ctx context.Context, pool *spiderpoolv2beta1.SpiderIPPool) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanAutoIPPoolLegacy checks whether the given IPPool should be deleted or not
func (ic *IPPoolController) cleanAutoIPPoolLegacy(ctx context.Context, pool *spiderpoolv2beta1.SpiderIPPool) error {
	_ = "STUB: not implemented"
	return nil
}

// check the label and decide to delete the IPPool or not

// unpack the IPPool corresponding application type,namespace and name

// check the IPPool's corresponding application whether is existed or not

// mismatch application UID
