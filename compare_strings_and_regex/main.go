package main

import "fmt"

func main() {
	var validSanitize int
	var validSanitizeRegex int
	var validSanitizeRegexWithoutQuantity int
	var validSanitizeLexical int
	var validSanitize3Phase int
	var validSanitizeNew int
	var validSanitizeBedTypes int
	for _, text := range data {
		sanitizedText := SanitizeManualPattern(text)
		if len(sanitizedText) > 0 {
			validSanitize += 1
		}
		sanitizedRegexText := SanitizeRegexPattern(text)
		if len(sanitizedRegexText) > 0 {
			validSanitizeRegex += 1
		}
		sanitizedRegexWithoutQuantityText := SanitizeRegexWithoutPattern(text)
		if len(sanitizedRegexWithoutQuantityText) > 0 {
			validSanitizeRegexWithoutQuantity += 1
		}

		sanitizedLexicalText := SanitizeRegexLexical(text)
		if len(sanitizedLexicalText) > 0 {
			validSanitizeLexical += 1
		}

		sanitized3PhaseText := Sanitize3Phase(text)
		if len(sanitized3PhaseText) > 0 {
			validSanitize3Phase += 1
		}

		sanitizedNewText := SanitizeComplexRegex(text)
		if len(sanitizedNewText) > 0 {
			validSanitizeNew += 1
		}

		sanitizedBedTypesText := SanitizeSimplifyLexical(text)
		if len(sanitizedBedTypesText) > 0 {
			validSanitizeBedTypes += 1
		}

		// if !isValidRegex(sanitizedRegexWithoutQuantityText, sanitizedLexicalText, sanitized3PhaseText) {
		// 	fmt.Println(sanitizedRegexWithoutQuantityText)
		// 	fmt.Println(sanitizedLexicalText)
		// 	fmt.Println(sanitized3PhaseText)
		// 	return
		// }
	}

	fmt.Printf("total valid SanitizeManualPattern:\t%d/%d\n", validSanitize, len(data))
	fmt.Printf("total valid SanitizeRegexPattern:\t%d/%d\n", validSanitizeRegex, len(data))
	fmt.Printf("total valid SanitizeRegexWithoutPattern:\t%d/%d\n", validSanitizeRegexWithoutQuantity, len(data))
	fmt.Printf("total valid SanitizeRegexLexical:\t%d/%d\n", validSanitizeLexical, len(data))
	fmt.Printf("total valid Sanitize3Phase:\t%d/%d\n", validSanitize3Phase, len(data))
	fmt.Printf("total valid SanitizeNew:\t%d/%d\n", validSanitizeNew, len(data))
	fmt.Printf("total valid SanitizeBedTypes:\t%d/%d\n", validSanitizeBedTypes, len(data))
}

func isValidRegex(t1, t2, t3 map[string]int) bool {
	for k, _ := range t1 {
		_, ok := t2[k]
		if !ok {
			return false
		}
		_, ok = t3[k]
		if !ok {
			return false
		}
	}
	return true
}
