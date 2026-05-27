// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationinformers

import (
	"k8s.io/client-go/tools/cache"
)

func (c *Controller) AddStatefulSetHandler(informer cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) onStatefulSetAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) onStatefulSetUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) onStatefulSetDelete(obj interface{}) { _ = "STUB: not implemented"; return }
