// Copyright 2023 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package cmd

import (
	"context"
)

func InitMultusDefaultCR(ctx context.Context, config *InitDefaultConfig, client *CoreClient) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchDefaultCNIName(defaultCNIName, cniDir string) (cniName, cniType string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
