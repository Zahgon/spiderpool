// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
)

type JobBehave string

const (
	JobTypeRunningForever JobBehave = "runningForeverJob"
	JobTypeFail           JobBehave = "failedJob"
	JobTypeFinish         JobBehave = "succeedJob"
)

func GenerateExampleJobYaml(behavior JobBehave, jdName, namespace string, parallelism *int32) *batchv1.Job {
	_ = "STUB: not implemented"
	return nil
}
