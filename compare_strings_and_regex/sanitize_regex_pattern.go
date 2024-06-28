package main

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	bedTypeRegex, _ = regexp.Compile("double|twin|single|king|semi-double|round|queen|sofa|tatami|kang|dorm|water|bunk")
)

func SanitizeRegexPattern(text string) map[string]int {
	if bedTypeRegex.MatchString(strings.ToLower(text)) {
		splittedText := strings.Split(text, " ")
		if len(splittedText) == 1 {
			return map[string]int{splittedText[0]: 1}
		}

		if len(splittedText) == 2 {
			if reflect.TypeOf(splittedText[0]).Kind() == reflect.String {
				return map[string]int{splittedText[0]: 1}
			}

			quantity, err := strconv.Atoi(splittedText[0])
			if err != nil {
				return nil
			}

			return map[string]int{splittedText[1]: quantity}
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

			if reflect.TypeOf(splittedText[1]).Kind() == reflect.String {
				sanitizedText[splittedText[1]] = q1
			}

			q2, err := strconv.Atoi(splittedText[3])
			if err != nil {
				return nil
			}

			if reflect.TypeOf(splittedText[4]).Kind() == reflect.String {
				sanitizedText[splittedText[4]] = q2
			}

			return sanitizedText
		}
	}

	return nil
}
