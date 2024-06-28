package main

import "testing"

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
