package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/tiket/TIX-AFFILIATE-COMMON-GO/collection"
	"golang.org/x/text/language"
)

var (
	data        = []string{"2 Double", "1 Double or Twin", "1 Double / 2 Twin", "1 Double", "1 Double or Twin", "2 Twin", "1 Twin", "1 Double", "1 Queen / 2 Single", "1 Queen / 2 Single", "3 Bunk Bed", "1 Double", "1 Twin", "1 Double", "2 Twin", "1 Double", "1 Double", "Double bed", "2 Twin", "1 Double", "1 Double", "1 Double", "2 Single", "1 Double", "1 King", "2 Single / 2 Bunk Bed", "2 Single", "1 King", "1 Double bed 131-150 width", "1 Queen-size bed 150-154 width", "1 Double bed 131-150 width", "1 Queen-size bed 150-154 width", "2 Double bed 131-150 width", "1 Queen-size bed 150-154 width", "1 Queen-size bed 150-154 width", "1 double bed", "2 single beds and 1 queen bed and 1 king bed", "1 double bed", "1 king bed and 1 queen bed", "1 double bed or 2 single beds", "2 Single", "1 Double", "1 Double", "2 Twin", "1 Double", "1 Double", "1 Extra Large Double", "1 King", "1 King", "1 Double", "1 Double", "1 Single & 1 Double", "2 Single", "1 Double", "3 Single & 1 Double", "1 Single & 1 Double & 1 Extra Large Double", "1 Double", "1 Extra Large Double", "1 Double & 1 Twin", "1 Double", "1 Double", "1 Double", "1 Double", "1 Double / 2 Twin", "2 Single / 1 Double", "2 Single / 1 Double", "1 Double", "1 Double", "2 Single / 1 Double", "1 Double / 2 Twin", "2 Single", "2 Single", "1 Double", "2 Twin", "2 Twin", "1 Double / 2 Single", "1 Double / 2 Single", "1 Double / 2 Single", "1 Double / 2 Single", "1 Double / 2 Single", "1 Double / 2 Single", "1 Double", "Double bed", "1 Double", "1 Double", "Double bed", "Double bed", "1 Double", "1 Single & 1 Double", "1 Double or Twin", "Double bed", "1 Double", "1 Double / 2 Single", "Double bed", "1 Double", "1 Twin", "1 Twin", "1 Double", "1 Double", "1 Double", "1 Double", "1 Double / 2 Twin", "1 Double / 2 Twin", "1 Double / 1 Twin", "2 Single bed 80-130 width", "1 double bed and 2 single beds", "1 double bed", "Double bed", "1 Bunk Bed", "1 Double", "Double bed", "1 Bunk Bed", "1 Double", "1 Bunk Bed", "Double bed", "Double bed", "Double bed", "1 King", "2 Twin", "1 king bed", "Bedroom 1: :1 king bed;Bedroom 2: :1 king bed or 1 king bed", "1 King-size bed 150-183 width", "1 Double", "1 King", "1 Bunk Bed", "1 Double", "Double bed", "Double", "1 Double", "1 King", "1 Twin / 1 King", "1 Twin / 1 King", "2 King", "1 Double", "1 King", "2 King", "1 Double", "Double bed", "Double bed", "1 King", "1 Single & 1 Double", "2 Single", "1 Double", "1 Single & 2 Double", "1 Double", "1 Double", "1 King", "1 King", "Double bed", "1 Double", "Single bed", "2 Twin", "Double bed", "1 Double", "1 King", "1 Double", "1 Double", "1 Double / 1 Twin", "Double bed", "1 Double / 1 Twin", "1 Double / 1 Twin", "1 double bed", "1 double bed", "1 Double or Twin", "2 single beds", "1 double bed", "1 double bed", "1 single bed", "1 King", "1 Double / 1 Twin", "1 King", "1 Double", "1 King", "1 Double", "1 Twin", "1 Double", "2 Twin", "4 King", "1 Double", "1 Double", "Twin bed", "1 Double", "1 Double", "1 Twin", "1 Double", "1 Double & 1 Twin", "Single bed", "Single bed", "1 Double", "1 Double", "Double bed", "Double bed", "1 Single & 1 Double", "1 Twin", "1 Single", "1 Single", "1 Double", "2 Single", "1 Double", "1 Twin", "1 Double", "1 King & 1 Bunk Bed", "2 Queen", "1 Single & 1 Double", "1 Double", "1 Double / 2 Single", "1 Queen / 2 Single", "1 Double", "1 Double", "2 Twin", "2 Twin", "2 Other", "2 Other", "1 Single & 1 Double", "1 Double", "1 Double & 2 Twin", "1 Double & 2 Twin", "1 Double & 2 Twin", "1 Double & 2 Twin", "1 Twin", "1 Twin", "1 King", "1 King", "1 Double", "1 King", "1 Double", "Double bed", "2 Single", "2 Twin", "1 Double", "1 Double", "1 Double / 1 Twin", "1 Twin", "1 Double", "1 Double", "1 Double", "2 Twin", "1 Double", "2 Twin", "2 Bunk Bed", "1 Twin / 2 Twin", "Large bed (King size)", "Large bed (King size)", "Large bed (King size)", "Double bed", "Twin bed", "Double bed", "1 Double", "Single bed", "Twin bed", "Double bed", "Single bed", "1 Double", "1 Double", "2 Double", "1 Twin", "1 Double", "1 Double", "1 Double", "3 King", "Double bed", "2 Double & 1 Twin", "2 Double & 1 Twin", "Double bed", "1 King", "1 Single & 1 Double", "1 Single & 2 King", "1 Double", "1 Double / 1 Twin", "1 Double / 1 Twin", "Double bed", "1 King", "2 Single / 1 Double", "Double bed", "1 Double", "Large bed (King size)", "Large bed (King size)", "1 Twin", "1 King", "1 Single", "1 King", "1 King", "1 Single & 1 Double", "1 Single & 1 Double", "1 Single & 1 Double", "1 Double", "1 Queen-size bed 150-154 width", "1 Queen-size bed 150-154 width", "1 Queen-size bed 150-154 width", "1 double bed or 2 single beds", "2 single beds and 1 king bed", "1 queen bed and 1 single bed", "1 double bed", "1 double bed and 1 single bed", "1 double bed", "1 king bed", "Bedroom 1: :1 king bed;Bedroom 2: :1 king bed;Common space :1 king bed", "2 king beds", "Bedroom 1: :1 king bed;Bedroom 2: :1 king bed or 2 single beds", "1 double bed", "Bedroom 1: :1 king bed;Bedroom 2: :1 king bed;Bedroom 3: :1 king bed or 2 single beds", "3 king beds", "5 king beds or 2 bunk beds", "1 King-size bed 150-183 width", "1 King-size bed 150-183 width", "2 single beds and 2 double beds", "1 king bed and 1 double bed", "2 double beds", "1 Double or Twin", "1 Double", "1 Double", "1 Single & 1 Double", "Double bed", "2 Twin", "1 Double or Twin", "Single bed", "Double bed", "Large bed (King size)", "Bedroom 1: :1 king bed;Bedroom 2: :1 single bed", "1 king bed", "1 king bed", "Bedroom 1: :1 king bed;Bedroom 2: :1 king bed", "2 Twin", "1 Double", "1 Double", "2 Twin", "1 Double", "Twin bed", "Twin bed", "Large bed (King size)", "Twin bed", "Twin bed", "Large bed (King size)", "2 Twin", "1 Double", "1 Double & 1 Twin", "2 Twin", "1 Double / 1 Twin", "2 Single & 1 Double", "1 Double", "1 Double", "2 Double", "1 Double", "1 Single & 1 Queen", "1 King", "1 King", "1 Double", "1 Double", "Twin bed", "Large bed (King size)", "Twin bed", "Single bed", "Large bed (King size)", "1 Double", "Twin bed", "1 Double", "2 Twin", "Single bed", "Twin bed", "1 Double / 2 Single", "1 Single & 1 Double", "2 Single / 1 Sofabed", "1 Single & 1 Double", "1 Double", "1 Double", "1 Twin", "Twin bed", "Single bed", "Single bed", "1 Twin", "1 King", "1 Double", "1 Double", "1 Single & 1 Double", "1 Double", "2 Double", "Large bed (King size)", "1 Twin", "1 Double", "1 Double", "2 Single & 1 King", "1 Double", "2 single beds or 1 double bed", "2 queen beds", "2 single beds or 1 double bed", "1 king bed", "1 king bed", "1 king bed or 2 single beds", "2 single beds or 1 double bed", "2 single beds or 1 double bed", "1 double bed and 1 single bed", "1 Double", "1 Double", "1 Double", "1 Twin", "1 King", "1 Double & 2 Twin", "Double bed", "Double bed", "Double bed", "1 Double", "1 Double / 2 Twin", "1 Double / 2 Twin", "1 Double / 2 Twin", "1 Double & 2 Twin", "2 Twin", "1 Double", "1 Single", "Single bed", "Single bed", "1 Double", "1 Single", "1 Single", "2 Bunk Bed", "1 Single", "1 Single", "1 Single & 1 Double", "1 Twin", "1 Bunk Bed", "Single bed", "1 Double", "Bunk bed", "Bunk bed", "1 Double", "2 Double", "Bunk bed", "3 Single", "1 Double", "1 Double", "2 Double", "2 Single", "1 Double", "2 Single", "1 Double", "1 Single & 1 Double", "1 Double", "2 Single", "1 Single", "Double bed", "Double bed", "Single bed", "1 Double", "Double bed", "1 Single", "1 Bunk Bed", "1 Double", "1 Single & 1 King", "1 Twin", "1 Double", "1 King", "2 Twin", "Single bed", "Bunk bed", "1 Single", "1 Single", "Single bed", "1 Queen", "Bunk bed", "Single bed", "Double bed", "Double bed", "1 King", "1 Single", "Double bed", "1 Double", "Single bed", "Single bed", "1 King", "1 Twin", "Single bed", "Double bed", "Single bed", "2 Single & 1 King", "Double bed", "1 King", "1 King-size bed 150-183 width", "1 King-size bed 150-183 width", "1 King-size bed 150-183 width", "2 Single bed 80-130 width"}
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

	bedTypeRegex, _      = regexp.Compile("double|twin|single|king|semi-double|round|queen|sofa|tatami|kang|dorm|water|bunk")
	quantityTypeRegex, _ = regexp.Compile(`\d+`)
	rdouble, _           = regexp.Compile("double")
	rtwin, _             = regexp.Compile("twin")
	rsingle, _           = regexp.Compile("single")
	rking, _             = regexp.Compile("king")
	rsemi, _             = regexp.Compile("semi-double")
	rround, _            = regexp.Compile("round")
	rqueen, _            = regexp.Compile("queen")
	rsofa, _             = regexp.Compile("sofa")
	rtatami, _           = regexp.Compile("tatami")
	rkang, _             = regexp.Compile("kang")
	rdorm, _             = regexp.Compile("dorm")
	rwater, _            = regexp.Compile("water")
	rbunk, _             = regexp.Compile("bunk")
	validRegexCollection = map[string]*regexp.Regexp{
		"double": rdouble,
		"twin":   rtwin,
		"single": rsingle,
		"king":   rking,
		"semi":   rsemi,
		"round":  rround,
		"queen":  rqueen,
		"sofa":   rsofa,
		"tatami": rtatami,
		"kang":   rkang,
		"dorm":   rdorm,
		"water":  rwater,
		"bunk":   rbunk,
	}

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

func checkValidWord(text string) bool {
	_, ok := bedTypesMap[text]
	return ok
}

func Sanitize(text string) map[string]int {
	text = strings.ToLower(text)
	splittedText := strings.Split(text, " ")
	if len(splittedText) == 1 {
		if ok := checkValidWord(splittedText[0]); ok {
			return map[string]int{splittedText[0]: 1}
		}
	}

	if len(splittedText) == 2 {
		if ok := checkValidWord(splittedText[0]); ok {
			return map[string]int{splittedText[0]: 1}
		}

		quantity, err := strconv.Atoi(splittedText[0])
		if err != nil {
			return nil
		}

		if ok := checkValidWord(splittedText[1]); ok {
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

		if ok := checkValidWord(splittedText[1]); ok {
			sanitizedText[splittedText[1]] = q1
		}

		q2, err := strconv.Atoi(splittedText[3])
		if err != nil {
			return nil
		}

		if ok := checkValidWord(splittedText[4]); ok {
			sanitizedText[splittedText[4]] = q2
		}

		return sanitizedText
	}

	return nil
}

func SanitizeRegex(text string) map[string]int {
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

func SanitizeRegexWithoutQuantity(text string) map[string]int {
	for key, reg := range validRegexCollection {
		if reg.MatchString(strings.ToLower(text)) {
			return map[string]int{key: 1}
		}
	}

	return nil
}

func SanitizeLexical(text string) map[string]int {
	// tokenization
	lexicals := getLexicals([]map[string]string{{language.English.String(): text}})
	// fmt.Println(lexicals)
	// get quantity map and bed type
	quantityMapByBedType := getQuantityMapByBedType(lexicals)

	return quantityMapByBedType
}

func SanitizeLexical2(text string) map[string]int {
	// tokenization
	lexicals := getLexicals2([]map[string]string{{language.English.String(): text}})
	// fmt.Println(lexicals)
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

func getLexicals2(bedTypeList []map[string]string) []lexical {
	var locLexical []lexical
	for _, bedTypeMapByLang := range bedTypeList {
		bedType := bedTypeMapByLang[language.English.String()]
		bedType = strings.ToLower(bedType)
		lexicals := parseBedTypeToLexical(bedType)
		locLexical = append(locLexical, lexicals...)
	}

	return locLexical
}

func parseBedTypeToLexical2(bedTypeText string) []lexical {
	var locLexical []lexical
	var tokenType tokenType
	splitBedType := strings.Split(bedTypeText, " ")
	for _, lexem := range splitBedType {
		if bedTypeRegex.MatchString(lexem) {
			tokenType = tokenBedType
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

func main() {
	var validSanitize int
	var validSanitizeRegex int
	var validSanitizeRegexWithoutQuantity int
	var validSanitizeLexical int
	// data = []string{"double bed"}
	for _, text := range data {
		// fmt.Println("TEXT = ", text)
		sanitizedText := Sanitize(text)
		if len(sanitizedText) > 0 {
			validSanitize += 1
		}
		sanitizedRegexText := SanitizeRegex(text)
		if len(sanitizedRegexText) > 0 {
			validSanitizeRegex += 1
		}
		sanitizedRegexWithoutQuantityText := SanitizeRegexWithoutQuantity(text)
		if len(sanitizedRegexWithoutQuantityText) > 0 {
			validSanitizeRegexWithoutQuantity += 1
		}

		sanitizedLexicalText := SanitizeLexical(text)
		if len(sanitizedLexicalText) > 0 {
			validSanitizeLexical += 1
		}

		for key, _ := range sanitizedText {
			lexKey, ok := sanitizedLexicalText[key]
			if !ok {
				fmt.Println("!!!WARNING!!!", lexKey, " ", key)
			}
		}
	}

	fmt.Printf("total valid Sanitize:%d/%d\n", validSanitize, len(data))
	fmt.Printf("total valid Sanitize:%d/%d\n", validSanitizeRegex, len(data))
	fmt.Printf("total valid Sanitize:%d/%d\n", validSanitizeRegexWithoutQuantity, len(data))
	fmt.Printf("total valid Sanitize:%d/%d\n", validSanitizeLexical, len(data))
}
