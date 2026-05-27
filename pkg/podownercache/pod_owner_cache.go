// Copyright 2024 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package podownercache

import (
	"context"

	"github.com/spidernet-io/spiderpool/pkg/lock"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PodOwnerCache struct {
	ctx       context.Context
	apiReader client.Reader

	cacheLock lock.RWMutex
	pods      map[types.NamespacedName]Pod
	ipToPod   map[string]types.NamespacedName
	// Cache for final owner references to reduce API calls, using pod NamespacedName as key
	ownerCache map[types.NamespacedName]*OwnerInfo
}

type Pod struct {
	types.NamespacedName
	OwnerInfo OwnerInfo
	IPs       []string
}

type OwnerInfo struct {
	APIVersion string
	Kind       string
	Namespace  string
	Name       string
}

type CacheInterface interface {
	GetPodByIP(ip string) *Pod
}

var logger *zap.Logger

func New(ctx context.Context, podInformer cache.SharedIndexInformer, apiReader client.Reader) (CacheInterface, error) {
	_ = "STUB: not implemented"
	return *new(CacheInterface), nil
}

func (s *PodOwnerCache) onPodAdd(obj any) { _ = "STUB: not implemented"; return }

func (s *PodOwnerCache) onPodUpdate(oldObj, newObj interface{}) { _ = "STUB: not implemented"; return }

func (s *PodOwnerCache) onPodDel(obj interface{}) { _ = "STUB: not implemented"; return }

func (s *PodOwnerCache) getFinalOwner(obj metav1.Object) (*OwnerInfo, error) {
	_ = "STUB: not implemented"
	return nil,

		// Create pod NamespacedName as the cache key
		nil
}

// Check if we already have a cached final owner for this pod

// If we found a cached result, return it immediately

// Assuming the first owner reference

// If not in cache, create the owner info

// Prepare an empty object of the owner kind

// Cache the negative result

// Set obj to the current owner to continue the loop

// Cache the final owner (or nil if no owner found)

func (s *PodOwnerCache) GetPodByIP(ip string) *Pod { _ = "STUB: not implemented"; return nil }
