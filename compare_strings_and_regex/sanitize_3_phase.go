package main

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	regConjuction, _ = regexp.Compile(`(or|and|&|/)`)
	regNumber, _     = regexp.Compile(`[1-9]+[0]?`)
)

func Sanitize3Phase(text string) map[string]int {
	result := map[string]int{}
	text = strings.ToLower(text)

	// phase 1 -> split by conjunction
	// 1 double and 1 double
	splitByConj := regConjuction.Split(text, -1)

	// phase 2 -> iterate and find valid bedType index
	// 1 double, 1 twin
	// rawQuantityMapByWord := map[string]string{}
	for _, rawWord := range splitByConj {
		// rawWord = 1 double
		// index [d]oubl[e] -> 1 double
		ind := regBedType.FindStringIndex(rawWord)
		if len(ind) == 2 {
			word := rawWord[ind[0]:ind[1]]
			qty := 1
			// phase 3 -> parse quantity
			quantity := regNumber.FindString(rawWord[:ind[0]]) // 1 d -> 1
			if intQuantity, err := strconv.Atoi(quantity); err == nil {
				qty = intQuantity
			}

			result[word] += qty
		}
	}

	return result
}
