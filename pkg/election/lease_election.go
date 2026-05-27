// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package election

import (
	"context"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection"

	"github.com/spidernet-io/spiderpool/pkg/lock"
)

var logger *zap.Logger

type SpiderLeaseElector interface {
	Run(ctx context.Context, clientSet kubernetes.Interface) error
	// IsElected returns a boolean value to check current Elector whether is a leader
	IsElected() bool
	GetLeader() string
}

type SpiderLeader struct {
	lock.RWMutex

	leaseLockName       string
	leaseLockNamespace  string
	leaseLockIdentity   string
	leaseDuration       time.Duration
	leaseRenewDeadline  time.Duration
	leaseRetryPeriod    time.Duration
	leaderRetryElectGap time.Duration

	isLeader      bool
	leaderElector *leaderelection.LeaderElector
}

// NewLeaseElector will return a SpiderLeaseElector object
func NewLeaseElector(leaseLockNS, leaseLockName, leaseLockIdentity string,
	leaseDuration, leaseRenewDeadline, leaseRetryPeriod, leaderRetryElectGap *time.Duration,
) (SpiderLeaseElector, error) {
	_ = "STUB: not implemented"
	return *new(SpiderLeaseElector), nil
}

func (sl *SpiderLeader) Run(ctx context.Context, clientSet kubernetes.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

// register will new client-go LeaderElector object with options configurations
func (sl *SpiderLeader) register(clientSet kubernetes.Interface) error {
	_ = "STUB: not implemented"
	return nil
}

// we can do cleanup here

func (sl *SpiderLeader) IsElected() bool { _ = "STUB: not implemented"; return false }

// tryToElect will elect continually
func (sl *SpiderLeader) tryToElect(ctx context.Context) { _ = "STUB: not implemented"; return }

// Once a node acquire the lease lock and become the leader, it will renew the lease lock continually until it failed to interact with API server.
// In this case the node will lose leader title and try to elect again.
// If there's a leader and another node will try to acquire the lease lock persistently until the leader renew failed.

func (sl *SpiderLeader) GetLeader() string { _ = "STUB: not implemented"; return "" }
