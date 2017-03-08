// Copyright 2017 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

import "github.com/hashicorp/golang-lru"

var (
	matchesCache    = createCache(100)
	topMatchesCache = createCache(10000)
)

func createCache(size int) *lru.Cache {
	c, err := lru.New(size)
	if err != nil {
		panic(err)
	}

	return c
}
