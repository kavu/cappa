// Copyright 2014 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

import (
	"testing"
)

const (
	uaString      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.76 Safari/537.36 OPR/16.0.1196.80"
	patternString = "Mozilla/5.0 (*Mac OS X 10_10*) AppleWebKit/* (KHTML, like Gecko)*Chrome/*Safari/*OPR/16.0*"
)

func TestNewPattern(t *testing.T) {
	_, err := NewPattern(patternString)
	if err != nil {
		t.Error(err)
	}
}

func TestMatch(t *testing.T) {
	p, err := NewPattern(patternString)
	if err != nil {
		t.Error(err)
	}

	if !p.Match(patternString) {
		t.Error("Pattern did not match")
	}
}

func BenchmarkNewPattern(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := NewPattern(patternString)
			if err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkMatch(b *testing.B) {
	p, err := NewPattern(patternString)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if !p.Match(uaString) {
				b.Error("Broken!")
			}
		}
	})
}
