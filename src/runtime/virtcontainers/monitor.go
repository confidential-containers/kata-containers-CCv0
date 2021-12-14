// Copyright (c) 2018 HyperHQ Inc.
//
// SPDX-License-Identifier: Apache-2.0
//

package virtcontainers

import (
	"context"
	"sync"
	"time"

	"github.com/kata-containers/kata-containers/src/runtime/virtcontainers/errors"
)

const (
	defaultCheckInterval = 5 * time.Second
	watcherChannelSize   = 128
)

// nolint: govet
type monitor struct {
	watchers []chan error
	sandbox  *Sandbox

	wg sync.WaitGroup
	sync.Mutex

	stopCh        chan bool
	checkInterval time.Duration

	running bool
}

func newMonitor(s *Sandbox) *monitor {
	return &monitor{
		sandbox:       s,
		checkInterval: defaultCheckInterval,
		stopCh:        make(chan bool, 1),
	}
}

func (m *monitor) newWatcher(ctx context.Context) (chan error, error) {
	m.Lock()
	defer m.Unlock()

	watcher := make(chan error, watcherChannelSize)
	m.watchers = append(m.watchers, watcher)

	if !m.running {
		m.running = true
		m.wg.Add(1)

		// create and start agent watcher
		go func() {
			tick := time.NewTicker(m.checkInterval)
			for {
				select {
				case <-m.stopCh:
					tick.Stop()
					m.wg.Done()
					return
				case <-tick.C:
					if err := m.watchHypervisor(ctx); err != nil {
						virtLog.Errorf("Monitor failed watching hypervisor: %s", err)
						m.wg.Done()
						return
					}
					if err := m.watchAgent(ctx); err != nil {
						virtLog.Errorf("Monitor failed watching agent: %s", err)
						m.wg.Done()
						return
					}
				}
			}
		}()
	}

	return watcher, nil
}

func (m *monitor) notify(ctx context.Context, err error) {
	errors.ErrorContext(&err, "Monitor marked the agent as dead")
	m.sandbox.agent.markDead(ctx, err)

	m.Lock()
	defer m.Unlock()

	if !m.running {
		return
	}

	// a watcher is not supposed to close the channel
	// but just in case...
	defer func() {
		if x := recover(); x != nil {
			virtLog.Warnf("watcher closed channel: %v", x)
		}
	}()

	for _, c := range m.watchers {
		// throw away message can not write to channel
		// make it not stuck, the first error is useful.
		select {
		case c <- err:

		default:
			virtLog.WithField("channel-size", watcherChannelSize).Warnf("watcher channel is full, throw notify message")
		}
	}
}

func (m *monitor) stop() {
	// wait outside of monitor lock for the watcher channel to exit.
	defer m.wg.Wait()

	m.Lock()
	defer m.Unlock()

	if !m.running {
		return
	}

	m.stopCh <- true
	defer func() {
		m.watchers = nil
		m.running = false
	}()

	// a watcher is not supposed to close the channel
	// but just in case...
	defer func() {
		if x := recover(); x != nil {
			virtLog.Warnf("watcher closed channel: %v", x)
		}
	}()

	for _, c := range m.watchers {
		close(c)
	}
}

func (m *monitor) watchAgent(ctx context.Context) (err error) {
	err = m.sandbox.agent.check(ctx)
	if err != nil {
		errors.ErrorContext(&err, "Monitor has failed to watch agent")
		m.notify(ctx, err)
	}
	return
}

func (m *monitor) watchHypervisor(ctx context.Context) (err error) {
	if err = m.sandbox.hypervisor.Check(); err != nil {
		errors.ErrorContext(&err, "Monitor has failed to watch hypervisor")
		m.notify(ctx, err)
	}
	return
}
