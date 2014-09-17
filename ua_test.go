package cappa

import (
	"testing"
)

func BenchmarkUAsMatches(b *testing.B) {
	patterns := ReadUAsFromCSV("test/browscap.csv")
	patterns.Matches(uaString)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if len(patterns.Matches(uaString)) < 1 {
				b.Error("Broken!")
			}
		}
	})
}

func BenchmarkUAsTopMatch(b *testing.B) {
	expected := "Mozilla/5.0 (*Mac OS X 10_10*) AppleWebKit/* (KHTML, like Gecko)*Chrome/*Safari/*OPR/16.0*"
	patterns := ReadUAsFromCSV("test/browscap.csv")
	patterns.Matches(uaString)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if patterns.TopMatch(uaString).PropertyName != expected {
				b.Error("Broken!")
			}
		}
	})
}
