package main

import (
	"regexp"
	"strings"
)

var (
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
)

func SanitizeRegexWithoutPattern(text string) map[string]int {
	for key, reg := range validRegexCollection {
		if reg.MatchString(strings.ToLower(text)) {
			return map[string]int{key: 1}
		}
	}

	return nil
}
