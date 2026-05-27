// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	frame "github.com/spidernet-io/e2eframework/framework"
	"sigs.k8s.io/controller-runtime/pkg/client"

	spiderpool "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
)

func CreateReservedIP(f *frame.Framework, ReservedIP *spiderpool.SpiderReservedIP, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to wait for finish last deleting

func DeleteReservedIPByName(f *frame.Framework, reservedIPName string, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func GetReservedIPByName(f *frame.Framework, reservedIPName string) *spiderpool.SpiderReservedIP {
	_ = "STUB: not implemented"
	return nil
}

func DeleteResverdIPUntilFinish(ctx context.Context, f *frame.Framework, reservedIPName string, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func GenerateExampleV4ReservedIPObject(ips []string) (string, *spiderpool.SpiderReservedIP) {
	_ = "STUB: not implemented"
	return "", nil
}

func GenerateExampleV6ReservedIPObject(ips []string) (string, *spiderpool.SpiderReservedIP) {
	_ = "STUB: not implemented"
	return "", nil
}
