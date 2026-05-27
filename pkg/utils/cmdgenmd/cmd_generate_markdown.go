// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmdgenmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var markdownPath string

// GenMarkDownCmd returns cobra.Command that help to generate markdown.
// The first param is the root cmd component name, the second one is the root cmd,
// the third one should be the root cmd logger.
func GenMarkDownCmd(component string, rootCmd *cobra.Command, logger *zap.Logger) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}
