package main

import (
	"strconv"
	"strings"

	"github.com/tiket/TIX-AFFILIATE-COMMON-GO/collection"
)

var (
	bedTypesMap = map[string]int{
		"double":      1,
		"twin":        1,
		"semi-double": 2,
		"king":        3,
		"round":       4,
		"queen":       5,
		"sofa":        6,
		"tatami":      7,
		"single":      9,
		"kang":        10,
		"space":       11,
		"capsule":     11,
		"dorm":        12,
		"water":       13,
		"bunk":        14,
	}
	bedTypesConjunction = collection.NewSet[string]("&", "/", "and", "or")
)

func SanitizeManualPattern(text string) map[string]int {
	text = strings.ToLower(text)
	splittedText := strings.Split(text, " ")
	if len(splittedText) == 1 {
		if isValidWord(splittedText[0]) {
			return map[string]int{splittedText[0]: 1}
		}
	}

	if len(splittedText) == 2 {
		if isValidWord(splittedText[0]) {
			return map[string]int{splittedText[0]: 1}
		}

		quantity, err := strconv.Atoi(splittedText[0])
		if err != nil {
			return nil
		}

		if isValidWord(splittedText[1]) {
			return map[string]int{splittedText[1]: quantity}
		}

		return nil
	}

	if len(splittedText) == 5 {
		if !bedTypesConjunction.Contains(splittedText[2]) {
			return nil
		}

		sanitizedText := map[string]int{}

		q1, err := strconv.Atoi(splittedText[0])
		if err != nil {
			return nil
		}

		if isValidWord(splittedText[1]) {
			sanitizedText[splittedText[1]] = q1
		}

		q2, err := strconv.Atoi(splittedText[3])
		if err != nil {
			return nil
		}

		if isValidWord(splittedText[4]) {
			sanitizedText[splittedText[4]] = q2
		}

		return sanitizedText
	}

	return nil
}

func isValidWord(text string) bool {
	_, ok := bedTypesMap[text]
	return ok
}
