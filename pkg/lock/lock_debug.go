// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build lockdebug
// +build lockdebug

package lock

import (
	"fmt"
	"io"
	"os"
	"time"

	deadlock "github.com/sasha-s/go-deadlock"

	"github.com/spidernet-io/spiderpool/pkg/logutils"
)

var (
	OutputWriter io.Writer = os.Stderr

	// SelfishThresholdSec is the number of seconds that should be used when
	// detecting if a lock was held for more than the specified time.
	SelfishThresholdSec = 0.5

	// Waiting for a lock for longer than DeadlockTimeout is considered a deadlock.
	// Ignored is DeadlockTimeout <= 0.
	DeadlockTimeout = 3 * time.Second
)

var (
	logger = logutils.Logger.Named("Debug-Lock")

	// selfishThresholdMsg is the message that will be printed when a lock was
	// held for more than selfishThresholdSec.
	selfishThresholdMsg = fmt.Sprintf("Goroutine took lock for more than %.2f seconds", SelfishThresholdSec)
)

func init() {
	deadlock.Opts.DeadlockTimeout = DeadlockTimeout
}

type internalRWMutex struct {
	deadlock.RWMutex
	t time.Time
}

func (i *internalRWMutex) Lock() { _ = "STUB: not implemented"; return }

func (i *internalRWMutex) Unlock() { _ = "STUB: not implemented"; return }

func (i *internalRWMutex) UnlockIgnoreTime() { _ = "STUB: not implemented"; return }

func (i *internalRWMutex) RLock() { _ = "STUB: not implemented"; return }

func (i *internalRWMutex) RUnlock() { _ = "STUB: not implemented"; return }

type internalMutex struct {
	deadlock.Mutex
	time.Time
}

func (i *internalMutex) Lock() { _ = "STUB: not implemented"; return }

func (i *internalMutex) Unlock() { _ = "STUB: not implemented"; return }

func (i *internalMutex) UnlockIgnoreTime() { _ = "STUB: not implemented"; return }

func printStackTo(sec float64, stack []byte, writer io.Writer) { _ = "STUB: not implemented"; return }

// A stack trace is usually in the following format:
// goroutine 1432 [running]:
// runtime/debug.Stack(0xc424c4a370, 0xc421f7f750, 0x1)
//   /usr/local/go/src/runtime/debug/stack.go:24 +0xa7
//   ...
// To know which trace belongs to which go routine we will append the
// go routine number to every line of the stack trace.

// Don't replace the last '\n'
