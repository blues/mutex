// Copyright 2022 Blues Inc.  All rights reserved.
// Use of this source code is governed by licenses granted by the
// copyright holder including that found in the LICENSE file.

package mutex

import (
	"sync"
)

type Mutex struct {
	Kind string
	m    sync.Mutex
}

func (m *Mutex) Lock() {
	m.m.Lock()
}

func (m *Mutex) Unlock() {
	m.m.Unlock()
}

type RWMutex struct {
	Kind string
	m    sync.RWMutex
}

func (m *RWMutex) Lock() {
	m.m.Lock()
}

func (m *RWMutex) Unlock() {
	m.m.Unlock()
}

func (m *RWMutex) RLock() {
	m.m.RLock()
}

func (m *RWMutex) RUnlock() {
	m.m.RUnlock()
}
