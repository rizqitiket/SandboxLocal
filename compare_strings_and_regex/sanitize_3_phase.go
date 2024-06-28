package main

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	rconjuction, _ = regexp.Compile(`(or|and|&|/)`)
	rnumber, _     = regexp.Compile(`[1-9][0-9]?`)
)

func Sanitize3Phase(text string) map[string]int {
	result := map[string]int{}
	text = strings.ToLower(text)

	// phase 1 -> split by conjunction
	splitByConj := rconjuction.Split(text, -1)

	// phase 2 -> iterate and find valid bedType index
	rawQuantityMapByWord := map[string]string{}
	for _, word := range splitByConj {
		ind := bedTypeRegex.FindStringIndex(word)
		if len(ind) == 2 {
			rawQuantityMapByWord[word] = word[:ind[0]]
			result[word[ind[0]:ind[1]]] = 0
		}
	}

	// phase 3 -> parse quantity
	for rawQuantity, word := range rawQuantityMapByWord {
		quantity := rnumber.FindString(rawQuantity)
		intQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			continue
		}
		result[word] += intQuantity
	}

	return result
}
