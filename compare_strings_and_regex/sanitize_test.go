package main

import "testing"

func BenchmarkSanitizer(b *testing.B) {
	var idx int

	b.Run("SanitizeManualPattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeManualPattern("1 Double or Twin")
		}
	})

	b.Run("SanitizeRegexPattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeRegexPattern("1 Double or Twin")
		}
	})

	b.Run("SanitizeRegexWithoutPattern", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeRegexWithoutPattern("1 Double or Twin")
		}
	})

	b.Run("SanitizeRegexLexical", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeRegexLexical("1 Double or Twin")
		}
	})

	b.Run("Sanitize3Phase", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Sanitize3Phase("1 Double or Twin")
		}
	})

	// ===========USING SAMPLE DATA===============

	b.Run("SanitizeManualPatternWithData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx = i % len(data)
			SanitizeManualPattern(data[idx])
		}
	})

	b.Run("SanitizeRegexPatternWithData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx = i % len(data)
			SanitizeRegexPattern(data[idx])
		}
	})

	b.Run("SanitizeRegexWithoutPatternWithData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx = i % len(data)
			SanitizeRegexWithoutPattern(data[idx])
		}
	})

	b.Run("SanitizeRegexLexicalWithData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx = i % len(data)
			SanitizeRegexLexical(data[idx])
		}
	})

	b.Run("Sanitize3PhaseWithData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			idx = i % len(data)
			Sanitize3Phase(data[idx])
		}
	})
}
