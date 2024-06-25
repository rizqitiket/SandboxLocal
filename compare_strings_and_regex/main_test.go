package main

import "testing"

var Idx int

func BenchmarkSanitizer(b *testing.B) {
	b.Run("Sanitize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Idx = i % len(Data)
			Sanitize(Data[Idx])
		}
	})

	b.Run("SanitizeRegex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Idx = i % len(Data)
			SanitizeRegex(Data[Idx])
		}
	})

}
