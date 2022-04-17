// Copyright 2022 Blues Inc.  All rights reserved.
// Use of this source code is governed by licenses granted by the
// copyright holder including that found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/blues/mutex"
)

var outerLock = mutex.Mutex{Type: "Outer"}
var innerLock = mutex.Mutex{Type: "Inner"}

func main() {
	outerLock.Lock()
	innerLock.Lock()
	innerLock.Unlock()
	outerLock.Unlock()
	innerLock.Lock()
	outerLock.Lock()
	fmt.Printf("failure to detect deadlock")
}
