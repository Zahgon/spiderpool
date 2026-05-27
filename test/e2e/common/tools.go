// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0
package common

import (
	"context"
	"math/rand"
	"net"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spidernet-io/spiderpool/pkg/types"
	corev1 "k8s.io/api/core/v1"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateString(lenNum int, isHex bool) string { _ = "STUB: not implemented"; return "" }

func GenerateRandomNumber(max int) string { _ = "STUB: not implemented"; return "" }

func CheckPodListInclude(list *corev1.PodList, pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func GetAdditionalPods(previous, latter *corev1.PodList) (pods []corev1.Pod) {
	_ = "STUB: not implemented"
	return nil
}

func ExecCommandOnKindNode(ctx context.Context, nodeNameList []string, command string) error {
	_ = "STUB: not implemented"
	return nil
}

func ExecCommand(ctx context.Context, cmd *exec.Cmd) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ContrastIpv6ToIntValues(ip1, ip2 string) error { _ = "STUB: not implemented"; return nil }

func SelectIPFromIps(version types.IPVersion, ips []net.IP, ipNum int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateRandomNumbers Given a number and a specified count of sub-numbers,
// generate the specified number of sub-numbers by randomly partitioning the given number.
func GenerateRandomNumbers(sum, user int) []int { _ = "STUB: not implemented"; return nil }
