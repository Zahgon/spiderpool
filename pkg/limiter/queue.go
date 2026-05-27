// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package limiter

import (
	"context"
	"errors"
	"sync"
)

type Limiter interface {
	AcquireTicket(ctx context.Context, tickets ...string) error
	ReleaseTicket(ctx context.Context, tickets ...string)
	Start(ctx context.Context) error
	Started() bool
}

func NewLimiter(c LimiterConfig) Limiter { _ = "STUB: not implemented"; return *new(Limiter) }

const DefaultTicket = "not to use this ticket"

var (
	ErrStartLimiteRrepeatedly = errors.New("start the limiter repeatedly")
	ErrShutdownQueue          = errors.New("queue shutdown")
	ErrFullQueue              = errors.New("queue is full")
)

type queue struct {
	cond           *sync.Cond
	shuttingDown   bool
	maxQueueSize   int
	elements       []*e
	grantedTickets map[string]int
}

type e struct {
	wantedTickets []string
	notifyCheckin chan empty
}

type empty struct{}

func (q *queue) AcquireTicket(ctx context.Context, tickets ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(iiiceoo): When ctx times out or is canceled, AcquireTicket should
// not still be blocked.

func (q *queue) queueUp(tickets ...string) (*e, error) { _ = "STUB: not implemented"; return nil, nil }

// When a new queuer begins to queue, here should try to wake up the
// conductor who may be rest in two cases at this time:
// 1. Queue is empty.
// 2. Checkin is blocking to avoid long polling.

func (q *queue) ReleaseTicket(ctx context.Context, tickets ...string) {
	_ = "STUB: not implemented"
	return
}

// When work is finished, the conductor who may be rest should be awakened
// to continue ticket checking. The reason for using Broadcast instead of
// Signal is that checkin and waitAllTicketsRetrieved will wait at the
// same time when the queue shutdown.

func (q *queue) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (q *queue) start() error { _ = "STUB: not implemented"; return nil }

func (q *queue) checkin() (shuttingDown bool) { _ = "STUB: not implemented"; return false }

// When no one is in queue, don't do meaningless ticket checking. Here may
// be awakened by the following cases:
// 1. A new queuer added.
// 2. An ongoing work has just been completed.
// 3. Queue shutdown.

// Waiting here for avoiding next unnecessary round of polling q.elements
// following cases could make it move on:
// 1. A new queuer call queueUp().
// 2. ReleaseTicket() when ticket revert.
// 3. shutdown() notify to close the queue.

func (q *queue) checkAvailableTicket(tickets ...string) bool {
	_ = "STUB: not implemented"
	return false
}

func (q *queue) grantTicket(e *e) { _ = "STUB: not implemented"; return }

func (q *queue) gracefulShutdown() { _ = "STUB: not implemented"; return }

func (q *queue) shutdown() { _ = "STUB: not implemented"; return }

// When the queue shutdown, notify the conductor do checkin once. If
// there are no queuers at this time, checkin successfully returns.
// Otherwise, after all queuers enter work, checkin returns.

func (q *queue) isAllTicketsRetrieved() bool { _ = "STUB: not implemented"; return false }

func (q *queue) waitAllTicketsRetrieved() { _ = "STUB: not implemented"; return }

// Make sure here don't wait for queue without working elements, as that
// could result in waiting for ReleaseTicket to be called on items in an
// empty queue which has already been shutdown, which will result in waiting
// indefinitely.

// Wait for a working elements to complete their work. Here will be awakened
// by ReleaseTicket when queue shutdown. When all the work to be completed
// is finished, gracefulShutdown will ensure that waitAllTicketsRetrieved will
// not be called again, and then shutdown safely.

func (q *queue) Started() bool { _ = "STUB: not implemented"; return false }
