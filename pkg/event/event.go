// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package event

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
)

// EventRecorder is Singleton
var EventRecorder record.EventRecorder

const FakeRecorderBufferSize = 1024

// init will give the EventRecorder with default fake Recorder to avoid panic if someone forget to initialize it
func init() {
	EventRecorder = record.NewFakeRecorder(FakeRecorderBufferSize)
}

// InitEventRecorder will initialize the Singleton EventRecorder
func InitEventRecorder(client *kubernetes.Clientset, scheme *runtime.Scheme, sourceComponent string) {
	_ = "STUB: not implemented"
	return
}
