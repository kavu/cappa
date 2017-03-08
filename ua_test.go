// Copyright 2017 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

import (
	"testing"
)

func BenchmarkUAsMatchesWithCacheWarm(b *testing.B) {
	patterns := ReadUAsFromCSV("test/browscap.csv")
	patterns.Matches(uaString)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			numMatches := len(patterns.Matches(uaString))

			if numMatches < 1 {
				b.Errorf("num matches is %d, expected to be =< 1", numMatches)
			}
		}
	})
}

func BenchmarkUAsTopMatchWithCacheWarm(b *testing.B) {
	patterns := ReadUAsFromCSV("test/browscap.csv")
	patterns.TopMatch(uaString)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			property := patterns.TopMatch(uaString).PropertyName
			if property != patternString {
				b.Errorf("expected %s, got %s", patternString, property)
			}
		}
	})
}
