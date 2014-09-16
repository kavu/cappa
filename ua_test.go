package cappa

import (
	"testing"
)

var (
	patterns = ReadUAsFromCSV("browscap_clean.csv")
)

func init() {
	patterns.Matches("")
}

func BenchmarkUAsMatch(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if len(patterns.Matches(uaString)) < 1 {
				b.Error("Broken!")
			}
		}
	})
}
