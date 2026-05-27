// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean resources",
	Long:  "Clean resources with specified parameters.",
	Run: func(cmd *cobra.Command, args []string) {
		validate, err := cmd.Flags().GetString("validate")
		if err != nil {
			logger.Fatal(err.Error())
			os.Exit(1)
		}
		mutating, err := cmd.Flags().GetString("mutating")
		if err != nil {
			logger.Fatal(err.Error())
			os.Exit(1)
		}
		logger.Sugar().Infof("validate %s\nmutating %s\n", validate, mutating)

		client, err := NewCoreClient()
		if err != nil {
			logger.Fatal(err.Error())
			os.Exit(1)
		}
		err = client.clean(validate, mutating)
		if err != nil {
			logger.Fatal(err.Error())
			os.Exit(1)
		}
	},
}

const (
	ENVNamespace          = "SPIDERPOOL_POD_NAMESPACE"
	ENVSpiderpoolInitName = "SPIDERPOOL_INIT_NAME"
)

type CoreClient struct {
	client.Client
}

func NewCoreClient() (*CoreClient, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *CoreClient) clean(validate, mutating string) error { _ = "STUB: not implemented"; return nil }

// Clean up MutatingWebhookConfiguration resources of spiderpool

// Clean up SriovNetworkResourcesInjectorMutating resources of sriov-network-operator

// Clean up sriov-operator-webhook-config resources of sriov-network-operator

// Clean up ValidatingWebhookConfiguration resources of spiderpool

// Clean up sriov-operator-webhook-config resources of sriov-network-operator

// Clean up SpiderIPPool resources of spiderpool

// Clean up SpiderSubnet resources of spiderpool

// Clean up SpiderEndpoint resources of spiderpool

// Clean up SpiderReservedIP resources of spiderpool

// Clean up SpiderMultusConfig resources of spiderpool

// Clean up SpiderCoordinator resources of spiderpool

// Delete all crds of spiderpool or sriov-network-operator

// Delete Job of spiderpool-Init

// cleanWebhookResources deletes a specific webhook configuration based on the provided resource type and name.
func (c *CoreClient) cleanWebhookResources(ctx context.Context, resourceType, resourceName string, obj client.Object) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanSpiderpoolResources lists and deletes specific Spiderpool resources, with an optional finalizer cleanup step.
func (c *CoreClient) cleanSpiderpoolResources(ctx context.Context, list client.ObjectList, resourceName string) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanCRDs lists and deletes CustomResourceDefinitions (CRDs) related to Spiderpool and sriov-network-operator.
func (c *CoreClient) cleanCRDs(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Delete all crds of sriov-network-operator
// After sriov-network-operator was uninstalled, sriov-network-operator did not delete its own CRD,
// and there were residual CRDs, which might bring some hidden dangers to the upgrade of sriov-network-operator;
// we tried to uninstall it through spiderpool.

// After helm uninstall, sriov-operator will delete the resources under sriovoperatorconfigs.sriovnetwork.openshift.io.
// If we delete this CRD resource in advance, helm uninstall will report an error.
// We will skip it for now to allow other resources to be deleted.

// cleanSpiderpoolInitJob deletes the spiderpool-init Job, logs any errors or success.
func (c *CoreClient) cleanSpiderpoolInitJob(ctx context.Context, spiderpoolInitNamespace, spiderpoolInitName string) error {
	_ = "STUB: not implemented"
	return nil
}
