package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	newRegex = regexp.MustCompile(`([1-9][0-9]?[ ]*)?(double|twin|single|king|semi\-double|round|queen|sofa|tatami|kang|dorm|water|bunk)`)
)

func findPrefixLastDigitIndex(s string) int {
	for idx, char := range s {
		if !unicode.IsDigit(char) {
			return idx - 1
		}
	}
	return -1
}

func SanitizeComplexRegex(text string) map[string]int {
	result := map[string]int{}
	text = strings.ToLower(text)

	rawBedTypes := newRegex.FindAllString(text, -1)
	for _, rawBedType := range rawBedTypes {
		//find last index of digit
		lastNumberIdx := findPrefixLastDigitIndex(rawBedType)

		bedType := rawBedType
		qty := 1

		if lastNumberIdx != -1 {
			qtyInt, err := strconv.Atoi(rawBedType[:lastNumberIdx+1])
			if err != nil {
				continue
			}
			qty = qtyInt
			bedType = strings.TrimSpace(rawBedType[lastNumberIdx+1:])
		}

		result[bedType] += qty
	}

	return result
}
