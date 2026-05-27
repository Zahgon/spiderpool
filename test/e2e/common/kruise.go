// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	kruisev1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	frame "github.com/spidernet-io/e2eframework/framework"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func GenerateExampleKruiseCloneSetYaml(name, namespace string, replica int32) *kruisev1.CloneSet {
	_ = "STUB: not implemented"
	return nil
}

func GenerateExampleKruiseStatefulSetYaml(name, namespace string, replica int32) *kruisev1.StatefulSet {
	_ = "STUB: not implemented"
	return nil
}

func CreateKruiseCloneSet(f *frame.Framework, kruiseCloneSet *kruisev1.CloneSet, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteKruiseCloneSetByName(f *frame.Framework, name, namespace string, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateKruiseStatefulSet(f *frame.Framework, kruiseStatefulSet *kruisev1.StatefulSet, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteKruiseStatefulSetByName(f *frame.Framework, name, namespace string, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func GetKruiseStatefulSet(f *frame.Framework, namespace, name string) (*kruisev1.StatefulSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ScaleKruiseStatefulSet(f *frame.Framework, kruiseStatefulSet *kruisev1.StatefulSet, replicas int32) (*kruisev1.StatefulSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
