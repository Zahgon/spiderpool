// Copyright 2025 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/containernetworking/cni/pkg/skel"
	"go.uber.org/zap"

	agentOpenAPIClient "github.com/spidernet-io/spiderpool/api/v1/agent/client"
)

// Set up file logging for spiderpool bin.
func SetupFileLogging(conf *NetConf) (*zap.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deleteIpamIps(spiderpoolAgentAPI *agentOpenAPIClient.SpiderpoolAgentAPI, args *skel.CmdArgs, k8sArgs K8sArgs) error {
	_ = "STUB: not implemented"
	return nil
}
