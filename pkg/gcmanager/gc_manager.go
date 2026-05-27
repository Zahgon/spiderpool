// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package gcmanager

import (
	"context"
	"time"

	"github.com/spidernet-io/spiderpool/pkg/election"
	"github.com/spidernet-io/spiderpool/pkg/ippoolmanager"
	"github.com/spidernet-io/spiderpool/pkg/kubevirtmanager"
	"github.com/spidernet-io/spiderpool/pkg/limiter"
	"github.com/spidernet-io/spiderpool/pkg/lock"
	"github.com/spidernet-io/spiderpool/pkg/nodemanager"
	"github.com/spidernet-io/spiderpool/pkg/podmanager"
	"github.com/spidernet-io/spiderpool/pkg/statefulsetmanager"
	"github.com/spidernet-io/spiderpool/pkg/workloadendpointmanager"

	"go.uber.org/zap"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

type GarbageCollectionConfig struct {
	EnableGCIP                                     bool
	EnableGCStatelessTerminatingPodOnReadyNode     bool
	EnableGCStatelessTerminatingPodOnNotReadyNode  bool
	EnableGCStatelessRunningPodOnEmptyPodStatusIPs bool
	EnableStatefulSet                              bool
	EnableKubevirtStaticIP                         bool
	EnableCleanOutdatedEndpoint                    bool

	ReleaseIPWorkerNum     int
	GCIPChannelBuffer      int
	MaxPodEntryDatabaseCap int
	WorkQueueMaxRetries    int

	DefaultGCIntervalDuration int
	TracePodGapDuration       int
	GCSignalTimeoutDuration   int
	GCSignalGapDuration       int
	AdditionalGraceDelay      int

	LeaderRetryElectGap time.Duration
}

var logger *zap.Logger

type GCManager interface {
	Start(ctx context.Context) <-chan error
	GetPodDatabase() PodDBer
	TriggerGCAll()
	Health() bool
}

var _ GCManager = &SpiderGC{}

type SpiderGC struct {
	k8ClientSet *kubernetes.Clientset
	PodDB       PodDBer

	// env configuration
	gcConfig *GarbageCollectionConfig

	// signal
	gcSignal         chan struct{}
	gcIPPoolIPSignal chan *PodEntry

	wepMgr      workloadendpointmanager.WorkloadEndpointManager
	ippoolMgr   ippoolmanager.IPPoolManager
	podMgr      podmanager.PodManager
	stsMgr      statefulsetmanager.StatefulSetManager
	kubevirtMgr kubevirtmanager.KubevirtManager
	nodeMgr     nodemanager.NodeManager
	leader      election.SpiderLeaseElector

	informerFactory informers.SharedInformerFactory
	gcLimiter       limiter.Limiter
	Locker          lock.Mutex
}

func NewGCManager(clientSet *kubernetes.Clientset, config *GarbageCollectionConfig,
	wepManager workloadendpointmanager.WorkloadEndpointManager,
	ippoolManager ippoolmanager.IPPoolManager,
	podManager podmanager.PodManager,
	stsManager statefulsetmanager.StatefulSetManager,
	kubevirtMgr kubevirtmanager.KubevirtManager,
	nodeMgr nodemanager.NodeManager,
	spiderControllerLeader election.SpiderLeaseElector,
) (GCManager, error) {
	_ = "STUB: not implemented"
	return *new(GCManager), nil
}

func (s *SpiderGC) Start(ctx context.Context) <-chan error { _ = "STUB: not implemented"; return nil }

// start pod informer

// trace pod worker

// monitor gc signal from CLI or DefaultGCInterval

func (s *SpiderGC) GetPodDatabase() PodDBer { _ = "STUB: not implemented"; return *new(PodDBer) }

func (s *SpiderGC) TriggerGCAll() { _ = "STUB: not implemented"; return }

const waitForCacheSyncTimeout = 5 * time.Second

func (s *SpiderGC) Health() bool { _ = "STUB: not implemented"; return false }
