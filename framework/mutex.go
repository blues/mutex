// Copyright 2022 Blues Inc.  All rights reserved.
// Use of this source code is governed by licenses granted by the
// copyright holder including that found in the LICENSE file.

package mutex

import (
	"sync"
)

// Mutex is the wrapper for both sync.Mutex and sync.RWMutex, used so that mutexes can
// declare their type.  Note that one difference between this package and sync is that
// the Write lock on RWMutex is called WLock/WUnlock rather than Lock/Unlock, so as
// to enable both simple and read-write mutexes to share this code.
type Mutex struct {
	mutex              sync.Mutex // Used for Lock/Unlock
	MutexInitialized   bool
	rwmutex            sync.RWMutex // Used for RLock/RUnlock/WLock/WUnlock
	RWmutexInitialized bool
	Type               string // The class of a mutex such as "customer database"
	Test               bool   // For activating testing behaviors in unit tests
	ID                 uint   // For use by an implementation to identify an instance
}

// By substituting your own callbacks, you can supply your own deadlock checking,
// determine how long the mutexes are being held, and so on.
type MutexCallbacks struct {
	Lock    func(m *Mutex)
	Unlock  func(m *Mutex)
	RLock   func(m *Mutex)
	RUnlock func(m *Mutex)
	WLock   func(m *Mutex)
	WUnlock func(m *Mutex)
}

// This is the current set of extensions
var mcb *MutexCallbacks

// Method by which mutex handlers can be inserted
func RegisterHandlers(mcbnew *MutexCallbacks) {
	mcb = mcbnew
}

// Default wrappers for sync package equivalents
func (m *Mutex) Lock() {
	if mcb == nil {
		m.mutex.Lock()
	} else {
		mcb.Lock(m)
	}
}

func (m *Mutex) Unlock() {
	if mcb == nil {
		m.mutex.Unlock()
	} else {
		mcb.Unlock(m)
	}
}

func (m *Mutex) WLock() {
	if mcb == nil {
		m.rwmutex.Lock()
	} else {
		mcb.WLock(m)
	}
}

func (m *Mutex) WUnlock() {
	if mcb == nil {
		m.rwmutex.Unlock()
	} else {
		mcb.WUnlock(m)
	}
}

func (m *Mutex) RLock() {
	if mcb == nil {
		m.rwmutex.RLock()
	} else {
		mcb.RLock(m)
	}
}

func (m *Mutex) RUnlock() {
	if mcb == nil {
		m.rwmutex.RUnlock()
	} else {
		mcb.RUnlock(m)
	}
}
