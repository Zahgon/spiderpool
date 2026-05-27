// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/containernetworking/cni/pkg/skel"
)

// CmdDel follows CNI SPEC cmdDel.
func CmdDel(args *skel.CmdArgs) (err error) { _ = "STUB: not implemented"; return nil }

// Defer a panic recover, so that in case we panic we can still return
// a proper error to the runtime.

// If it is recovering and an error occurs, then we need to
// present both.
