// Copyright 2017 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

import (
	"testing"
)

const (
	uaString      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
	patternString = "Mozilla/5.0 (*Mac OS X 10?12*) applewebkit* (*khtml*like*gecko*) Chrome/56.*Safari/*"
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

	if !p.Match(uaString) {
		t.Errorf("pattern %s did not match agent %s", patternString, uaString)
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
