package main

import "testing"

var idx int

func BenchmarkSanitizer(b *testing.B) {
	b.Run("Sanitize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Sanitize("1 Double or Twin")
		}
	})

	b.Run("SanitizeRegex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeRegex("1 Double or Twin")
		}
	})

	b.Run("SanitizeRegexWithoutQuanity", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeRegexWithoutQuantity("1 Double or Twin")
		}
	})

	b.Run("SanitizeWithModulus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx = i % len(data)
			Sanitize(data[idx])
		}
	})

	b.Run("SanitizeRegexWithModulus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx = i % len(data)
			SanitizeRegex(data[idx])
		}
	})

	b.Run("SanitizeRegexWithoutQuantityWithModulus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeRegexWithoutQuantity(data[idx])
		}
	})
}
