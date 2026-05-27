// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationinformers

import (
	"k8s.io/client-go/tools/cache"
)

func (c *Controller) AddDaemonSetHandler(informer cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) onDaemonSetAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) onDaemonSetUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) onDaemonSetDelete(obj interface{}) { _ = "STUB: not implemented"; return }
