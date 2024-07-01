package main

import (
	"fmt"
	"testing"
)

func BenchmarkSanitizer(b *testing.B) {
	// b.Run("SanitizeManualPattern", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		SanitizeManualPattern("1 Double or Twin")
	// 	}
	// })

	// b.Run("SanitizeRegexPattern", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		SanitizeRegexPattern("1 Double or Twin")
	// 	}
	// })

	// b.Run("SanitizeRegexWithoutPattern", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		SanitizeRegexWithoutPattern("1 Double or Twin")
	// 	}
	// })

	// b.Run("SanitizeRegexLexical", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		SanitizeRegexLexical("1 Double or Twin")
	// 	}
	// })

	// b.Run("Sanitize3Phase", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		Sanitize3Phase("1 Double or Twin")
	// 	}
	// })

	// ===========USING SAMPLE DATA===============

	b.Run("SanitizeManualPatternWithData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, d := range data {
				SanitizeManualPattern(d)
			}
		}
	})

	// b.Run("SanitizeRegexPatternWithData", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		for _, d := range data {
	// 			SanitizeRegexPattern(d)
	// 		}
	// 	}

	// })

	// b.Run("SanitizeRegexWithoutPatternWithData", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		for _, d := range data {
	// 			SanitizeRegexWithoutPattern(d)
	// 		}
	// 	}
	// })

	b.Run("SanitizeOneComplexRegex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, d := range data {
				SanitizeComplexRegex(d)
			}
		}
	})

	b.Run("SanitizeRegexLexical", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, d := range data {
				SanitizeRegexLexical(d)
			}
		}
	})

	// b.Run("Sanitize3PhaseWithData", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		for _, d := range data {
	// 			Sanitize3Phase(d)
	// 		}
	// 	}
	// })

	b.Run("SanitizeSimplifyRegexLexical", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, d := range data {
				SanitizeSimplifyLexical(d)
			}
		}
	})
}

func TestXxx(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		result := SanitizeComplexRegex("Bedroom 1: :1 king bed;Bedroom 2: :1 king bed or 1 king bed")
		fmt.Println(result)
		result = SanitizeRegexLexical("Bedroom 1: :1 king bed;Bedroom 2: :1 king bed or 1 king bed")
		fmt.Println(result)
		result = SanitizeSimplifyLexical("1Bedroom 1: :1 king bed;Bedroom 2: :1 king bed or 1 king bed")
		fmt.Println(result)
	})
}
