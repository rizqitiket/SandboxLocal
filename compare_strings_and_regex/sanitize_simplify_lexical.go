package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	regQty      = regexp.MustCompile(`[1-9][0-9]*`)
	empty       struct{}
	bedTypeDict = map[string]struct{}{
		"double":      empty,
		"twin":        empty,
		"single":      empty,
		"king":        empty,
		"semi-double": empty,
		"round":       empty,
		"queen":       empty,
		"sofa":        empty,
		"tatami":      empty,
		"kang":        empty,
		"dorm":        empty,
		"water":       empty,
		"bunk":        empty,
	}
)

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, c := range s[1:] {
		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}

func SanitizeSimplifyLexical(text string) map[string]int {
	result := map[string]int{}
	text = strings.ToLower(text)
	words := strings.FieldsFunc(text, func(char rune) bool {
		return char == ' ' || char == '-' || char == '(' || char == ')'
	})

	lastLookupQty := 1
	for _, word := range words {
		if _, ok := bedTypeDict[word]; ok {
			result[word] += lastLookupQty
			lastLookupQty = 1
			continue
		}

		if isNumber(word) {
			if qty, err := strconv.Atoi(word); err == nil {
				lastLookupQty = qty
			}
		}
	}

	return result
}
