// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(spiderpoolv2beta1.AddToScheme(scheme))
}

const retryIntervalSec = 2

type CoreClient struct {
	client.Client
}

func NewCoreClient() (*CoreClient, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *CoreClient) WaitForCoordinatorCreated(ctx context.Context, coord *spiderpoolv2beta1.SpiderCoordinator) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CoreClient) WaitForSubnetCreated(ctx context.Context, subnet *spiderpoolv2beta1.SpiderSubnet) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CoreClient) WaitForIPPoolCreated(ctx context.Context, ipPool *spiderpoolv2beta1.SpiderIPPool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CoreClient) WaitForEndpointReady(ctx context.Context, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CoreClient) CheckEndpointsAvailable(ctx context.Context, namespace, name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CoreClient) WaitMultusCNIConfigCreated(ctx context.Context, multuscniconfig *spiderpoolv2beta1.SpiderMultusConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CoreClient) WaitPodListReady(ctx context.Context, namespace string, labels map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}
