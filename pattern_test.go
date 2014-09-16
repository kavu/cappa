package cappa

import (
	"testing"
)

const (
	uaString      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.94 Safari/537.36"
	patternString = "Mozilla/5.0 (*Mac OS X 10_9*) AppleWebKit/* (KHTML, like Gecko)*Chrome/37.*Safari/*"
)

func BenchmarkPattrenNewPattern(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := NewPattern(patternString)
			if err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkPattrenMatchUA(b *testing.B) {
	p, err := NewPattern(patternString)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if !p.MatchUA(uaString) {
				b.Error("Broken!")
			}
		}
	})
}
