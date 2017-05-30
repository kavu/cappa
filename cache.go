// Copyright 2017 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

import "github.com/hashicorp/golang-lru"

// Cache describes interface for effective UA results matching.
// You can implement your own Cache if you wish and set them
// to MatchesCache and TopMatchesCache.
type Cache interface {
	Get(key interface{}) (interface{}, bool)
	Add(key, value interface{}) bool
}

// MatchesCache stores a all matched UAs for single uniqueUser-Agent string.
// By default it uses the LRU cache of 100 elements
// created with NewDefaultCache.
var MatchesCache Cache = NewDefaultCache(100)

// TopMatchesCache stores a top matched UA for single unique User-Agent string.
// By default it uses the LRU cache of 100000 elements
// created with NewDefaultCache.
var TopMatchesCache Cache = NewDefaultCache(10000)

// NewDefaultCache returns a fixed size LRU cache backed by
// https://github.com/hashicorp/golang-lru which implements the Cache
func NewDefaultCache(size int) *lru.Cache {
	c, err := lru.New(size)
	if err != nil {
		panic(err)
	}

	return c
}
