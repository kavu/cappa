// Copyright 2017 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

import (
	"testing"
)

func TestReadUAsFromCSV(t *testing.T) {
	patterns := ReadUAsFromCSV("test/browscap.csv")

	if len(patterns) < 1 {
		t.Errorf("num UAs is %d, expected to be >= 1", len(patterns))
	}
}

func TestMatches(t *testing.T) {
	patterns := ReadUAsFromCSV("test/browscap.csv")
	matches := patterns.Matches(uaString)

	if len(matches) < 1 {
		t.Errorf("num matches is %d, expected to be =< 1", len(matches))
	}
}

func TestTopMatch(t *testing.T) {
	patterns := ReadUAsFromCSV("test/browscap.csv")
	match := patterns.TopMatch(uaString)

	if match != nil {
		property := match.PropertyName

		if property != patternString {
			t.Errorf("expected %s, got %s", patternString, property)
		}
	} else {
		t.Errorf("expected match to be non-nil")

	}
}

func BenchmarkReadUAsFromCSV(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			patterns := ReadUAsFromCSV("test/browscap.csv")

			if len(patterns) < 1 {
				b.Errorf("num UAs is %d, expected to be >= 1", len(patterns))
			}
		}
	})
}

func BenchmarkMatchesWithCacheWarm(b *testing.B) {
	patterns := ReadUAsFromCSV("test/browscap.csv")
	patterns.Matches(uaString)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			matches := patterns.Matches(uaString)

			if len(matches) < 1 {
				b.Errorf("num matches is %d, expected to be >= 1", len(matches))
			}
		}
	})
}

func BenchmarkTopMatchWithCacheWarm(b *testing.B) {
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
