package main

import (
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/language"
)

var (
	quantityTypeRegex, _ = regexp.Compile(`\d+`)

	tokenQuantity = tokenType("QUANTITY")
	tokenBedType  = tokenType("BEDTYPE")
)

type (
	lexical struct {
		lexemes string
		token   tokenType
	}

	tokenType string
)

func SanitizeRegexLexical(text string) map[string]int {
	// tokenization
	lexicals := getLexicals([]map[string]string{{language.English.String(): text}})
	// get quantity map and bed type
	quantityMapByBedType := getQuantityMapByBedType(lexicals)

	return quantityMapByBedType
}

func getQuantityMapByBedType(lexicals []lexical) map[string]int {
	result := map[string]int{}
	if len(lexicals) == 1 && lexicals[0].token == tokenBedType {
		return map[string]int{lexicals[0].lexemes: 1}
	}

	for i := 0; i < len(lexicals)-1; i++ {
		if lexicals[i].token == tokenQuantity && lexicals[i+1].token == tokenBedType {
			quantity, err := strconv.Atoi(lexicals[i].lexemes)
			if err != nil {
				continue
			}
			result[lexicals[i+1].lexemes] += quantity
			i += 1
		} else if lexicals[i].token == tokenBedType {
			result[lexicals[i+1].lexemes] += 1
		}
	}

	return result
}

func getLexicals(bedTypeList []map[string]string) []lexical {
	var locLexical []lexical
	for _, bedTypeMapByLang := range bedTypeList {
		bedType := bedTypeMapByLang[language.English.String()]
		bedType = strings.ToLower(bedType)
		lexicals := parseBedTypeToLexical(bedType)
		locLexical = append(locLexical, lexicals...)
	}

	return locLexical
}

func parseBedTypeToLexical(bedTypeText string) []lexical {
	var locLexical []lexical
	var tokenType tokenType
	splitBedType := strings.Split(bedTypeText, " ")
	for _, lexem := range splitBedType {
		if bedTypeRegex.MatchString(lexem) {
			tokenType = tokenBedType
		} else if quantityTypeRegex.MatchString(lexem) {
			tokenType = tokenQuantity
		} else if checkStringInt(lexem) {
			tokenType = tokenQuantity
		} else {
			continue
		}
		locLexical = append(locLexical, lexical{
			lexemes: lexem,
			token:   tokenType,
		})
	}

	return locLexical
}

func checkStringInt(text string) bool {
	if _, err := strconv.Atoi(text); err != nil {
		return false
	}
	return true
}
