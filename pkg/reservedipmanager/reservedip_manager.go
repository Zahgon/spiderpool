// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package reservedipmanager

import (
	"context"
	"net"

	"sigs.k8s.io/controller-runtime/pkg/client"

	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	"github.com/spidernet-io/spiderpool/pkg/types"
)

type ReservedIPManager interface {
	GetReservedIPByName(ctx context.Context, rIPName string, cached bool) (*spiderpoolv2beta1.SpiderReservedIP, error)
	ListReservedIPs(ctx context.Context, cached bool, opts ...client.ListOption) (*spiderpoolv2beta1.SpiderReservedIPList, error)
	AssembleReservedIPs(ctx context.Context, version types.IPVersion) ([]net.IP, error)
}

type reservedIPManager struct {
	client    client.Client
	apiReader client.Reader
}

func NewReservedIPManager(client client.Client, apiReader client.Reader) (ReservedIPManager, error) {
	_ = "STUB: not implemented"
	return *new(ReservedIPManager), nil
}

func (rm *reservedIPManager) GetReservedIPByName(ctx context.Context, rIPName string, cached bool) (*spiderpoolv2beta1.SpiderReservedIP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *reservedIPManager) ListReservedIPs(ctx context.Context, cached bool, opts ...client.ListOption) (*spiderpoolv2beta1.SpiderReservedIPList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *reservedIPManager) AssembleReservedIPs(ctx context.Context, version types.IPVersion) ([]net.IP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
