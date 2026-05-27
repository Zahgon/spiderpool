// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationinformers

import (
	"k8s.io/client-go/tools/cache"
)

func (c *Controller) AddReplicaSetHandler(informer cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) onReplicaSetAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) onReplicaSetUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) onReplicaSetDelete(obj interface{}) { _ = "STUB: not implemented"; return }
