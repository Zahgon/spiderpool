// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationinformers

import (
	"k8s.io/client-go/tools/cache"
)

func (c *Controller) AddJobController(informer cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) onJobAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) onJobUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) onJobDelete(obj interface{}) { _ = "STUB: not implemented"; return }
