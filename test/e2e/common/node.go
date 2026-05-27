// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	e2e "github.com/spidernet-io/e2eframework/framework"
)

// Restart the node and wait for the cluster to be ready and the Pods in the cluster to "running".
// In a "Kind" cluster, it is not recommended to set `nodes` to `nil`.
// Restarting all nodes will cause the spiderpool component to fail to pull up, further rendering the cluster unavailable.
func RestartNodeUntilClusterReady(ctx context.Context, frame *e2e.Framework, nodes ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Waiting for nodes to be ready

func GetNodeNetworkInfo(ctx context.Context, frame *e2e.Framework, nodeList []string) error {
	_ = "STUB: not implemented"
	return nil
}
