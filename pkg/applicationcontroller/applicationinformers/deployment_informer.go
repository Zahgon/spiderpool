// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationinformers

import (
	"k8s.io/client-go/tools/cache"
)

func (c *Controller) AddDeploymentHandler(informer cache.SharedIndexInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) onDeploymentAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) onDeploymentUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) onDeploymentDelete(obj interface{}) { _ = "STUB: not implemented"; return }
