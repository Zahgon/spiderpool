// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package applicationinformers

import (
	"context"

	"go.uber.org/zap"
)

var controllersLogger *zap.Logger

type (
	AppInformersAddOrUpdateFunc func(ctx context.Context, oldObj, newObj interface{}) error
	APPInformersDelFunc         func(ctx context.Context, obj interface{}) error
)

type Controller struct {
	reconcileFunc AppInformersAddOrUpdateFunc
	cleanupFunc   APPInformersDelFunc
}

func NewApplicationController(reconcile AppInformersAddOrUpdateFunc, cleanup APPInformersDelFunc, logger *zap.Logger) (*Controller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
