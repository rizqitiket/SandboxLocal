package main

import "testing"

func BenchmarkSanitizer(b *testing.B) {
	var idx int

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

	// b.Run("SanitizeRegexWithoutQuanity", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		SanitizeRegexWithoutQuantity("1 Double or Twin")
	// 	}
	// })

	b.Run("SanitizeLexical", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeLexical("1 Double or Twin")
		}
	})

	b.Run("SanitizeLexical2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeLexical2("1 Double or Twin")
		}
	})

	// ====================

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

	// b.Run("SanitizeRegexWithoutQuantityWithModulus", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		SanitizeRegexWithoutQuantity(data[idx])
	// 	}
	// })

	b.Run("SanitizeLexicalWithModulus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeLexical(data[idx])
		}
	})

	b.Run("SanitizeLexical2WithModulus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SanitizeLexical2(data[idx])
		}
	})
}

func BenchmarkNewRoomTPAExtension(b *testing.B) {
	var bedTypeList = []map[string]string{
		{
			"en": "1 Double or Twin",
		},
		{
			"en": "1 Double / 2 Twin",
		},
		{
			"en": "1 Double",
		},
		{
			"en": "1 asd",
		},
		{
			"en": "1 re",
		},
		{
			"en": "1 ggefg",
		},
		{
			"en": "1 asd",
		},
		{
			"en": "1 re",
		},
		{
			"en": "1 ggefg",
		},
		{
			"en": "1 asd",
		},
		{
			"en": "1 re",
		},
		{
			"en": "1 ggefg",
		},
	}

	b.Run("SanitizeRoomBedType", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewRoomTPAExtension(bedTypeList)
		}
	})

	b.Run("SanitizeRoomBedTypeCombine", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewRoomTPAExtensionCombine(bedTypeList)
		}
	})

	b.Run("SanitizeRoomBedTypeRegex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewRoomTPAExtensionRegex(bedTypeList)
		}
	})

	b.Run("SanitizeRoomBedTypeRegex2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewRoomTPAExtensionRegex2(bedTypeList)
		}
	})
}
