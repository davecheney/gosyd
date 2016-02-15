// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	// reduce this number and see the effect
	const workers = 99

	var wg sync.WaitGroup
	wg.Add(workers)
	m := map[int]int{}
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				m[i]++ // HL
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}

// END OMIT
