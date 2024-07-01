package main

import (
	"reflect"
	"strconv"
	"strings"

	"log/slog"
)

type (
	defaultRoomTPAExtension struct {
		BedTypes defaultBedTypes `xml:"ns4:BedTypes"`
	}

	defaultBedTypes struct {
		BedType []defaultBedType `xml:"ns4:BedType"`
	}

	defaultBedType struct {
		BedTypeCode  int     `xml:"BedTypeCode,attr,omitempty"`
		Quantity     int     `xml:"Quantity,attr,omitempty"`
		Length       float64 `xml:"Length,attr,omitempty"`
		Width        float64 `xml:"Width,attr,omitempty"`
		CategoryCode int     `xml:"CategoryCode,attr,omitempty"`
	}
)

func NewRoomTPAExtension(bedTypeList []map[string]string) *defaultRoomTPAExtension {
	quantityMapByBedType := map[string]int{}
	for _, bedType := range bedTypeList {
		lowerText := strings.ToLower(bedType["en"])
		// strings.Split guaranteed to return slice with length at least 1
		splittedText := strings.Split(lowerText, " ")
		if _, ok := bedTypesMap[splittedText[0]]; ok && len(splittedText) <= 2 {
			quantityMapByBedType[splittedText[0]] += 1
		} else if len(splittedText) == 2 {
			quantity, err := strconv.Atoi(splittedText[0])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if _, ok := bedTypesMap[splittedText[1]]; ok {
				quantityMapByBedType[splittedText[1]] += quantity
			}
		} else if len(splittedText) == 5 && bedTypesConjunction.Contains(splittedText[2]) {
			quantity, err := strconv.Atoi(splittedText[0])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if _, ok := bedTypesMap[splittedText[1]]; ok {
				quantityMapByBedType[splittedText[1]] += quantity
			}

			quantity, err = strconv.Atoi(splittedText[3])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if _, ok := bedTypesMap[splittedText[4]]; ok {
				quantityMapByBedType[splittedText[4]] += quantity
			}
		}

	}
	var bedTypes []defaultBedType
	for bedType, quantity := range quantityMapByBedType {
		bedTypes = append(bedTypes, defaultBedType{
			BedTypeCode: bedTypesMap[bedType],
			Quantity:    quantity,
		})
	}
	return &defaultRoomTPAExtension{
		BedTypes: defaultBedTypes{
			BedType: bedTypes,
		},
	}
}

func NewRoomTPAExtensionCombine(bedTypeList []map[string]string) *defaultRoomTPAExtension {
	quantityMapByBedType := map[string]int{}
	for _, bedType := range bedTypeList {
		lowerText := strings.ToLower(bedType["en"])
		if !regBedType.MatchString(lowerText) {
			continue
		}
		// strings.Split guaranteed to return slice with length at least 1
		splittedText := strings.Split(lowerText, " ")
		if _, ok := bedTypesMap[splittedText[0]]; ok && len(splittedText) <= 2 {
			quantityMapByBedType[splittedText[0]] += 1
		} else if len(splittedText) == 2 {
			quantity, err := strconv.Atoi(splittedText[0])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if _, ok := bedTypesMap[splittedText[1]]; ok {
				quantityMapByBedType[splittedText[1]] += quantity
			}
		} else if len(splittedText) == 5 && bedTypesConjunction.Contains(splittedText[2]) {
			quantity, err := strconv.Atoi(splittedText[0])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if _, ok := bedTypesMap[splittedText[1]]; ok {
				quantityMapByBedType[splittedText[1]] += quantity
			}

			quantity, err = strconv.Atoi(splittedText[3])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if _, ok := bedTypesMap[splittedText[4]]; ok {
				quantityMapByBedType[splittedText[4]] += quantity
			}
		}

	}
	var bedTypes []defaultBedType
	for bedType, quantity := range quantityMapByBedType {
		bedTypes = append(bedTypes, defaultBedType{
			BedTypeCode: bedTypesMap[bedType],
			Quantity:    quantity,
		})
	}
	return &defaultRoomTPAExtension{
		BedTypes: defaultBedTypes{
			BedType: bedTypes,
		},
	}
}

func NewRoomTPAExtensionRegex(bedTypeList []map[string]string) *defaultRoomTPAExtension {
	quantityMapByBedType := map[string]int{}
	for _, bedType := range bedTypeList {
		lowerText := strings.ToLower(bedType["en"])
		if !regBedType.MatchString(lowerText) {
			continue
		}
		// strings.Split guaranteed to return slice with length at least 1
		splittedText := strings.Split(lowerText, " ")
		if reflect.TypeOf(splittedText[0]).Kind() == reflect.String && len(splittedText) <= 2 {
			quantityMapByBedType[splittedText[0]] += 1
		} else if len(splittedText) == 2 {
			quantity, err := strconv.Atoi(splittedText[0])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if reflect.TypeOf(splittedText[1]).Kind() == reflect.String {
				quantityMapByBedType[splittedText[1]] += quantity
			}
		} else if len(splittedText) == 5 && bedTypesConjunction.Contains(splittedText[2]) {
			quantity, err := strconv.Atoi(splittedText[0])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if reflect.TypeOf(splittedText[1]).Kind() == reflect.String {
				quantityMapByBedType[splittedText[1]] += quantity
			}

			quantity, err = strconv.Atoi(splittedText[3])
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if reflect.TypeOf(splittedText[4]).Kind() == reflect.String {
				quantityMapByBedType[splittedText[4]] += quantity
			}
		}

	}
	var bedTypes []defaultBedType
	for bedType, quantity := range quantityMapByBedType {
		bedTypes = append(bedTypes, defaultBedType{
			BedTypeCode: bedTypesMap[bedType],
			Quantity:    quantity,
		})
	}
	return &defaultRoomTPAExtension{
		BedTypes: defaultBedTypes{
			BedType: bedTypes,
		},
	}
}

func NewRoomTPAExtensionRegex2(bedTypeList []map[string]string) *defaultRoomTPAExtension {
	quantityMapByBedType := map[string]int{}
	for _, bedType := range bedTypeList {
		lowerText := strings.ToLower(bedType["en"])
		if !regBedType.MatchString(lowerText) {
			continue
		}

		for key, reg := range validRegexCollection {
			if reg.MatchString(lowerText) {
				quantityMapByBedType[key] = 1
			}
		}
	}
	var bedTypes []defaultBedType
	for bedType, quantity := range quantityMapByBedType {
		bedTypes = append(bedTypes, defaultBedType{
			BedTypeCode: bedTypesMap[bedType],
			Quantity:    quantity,
		})
	}
	return &defaultRoomTPAExtension{
		BedTypes: defaultBedTypes{
			BedType: bedTypes,
		},
	}
}

func NewRoomTPAExtensionIndex(bedTypeList []map[string]string) *defaultRoomTPAExtension {
	quantityMapByBedType := map[string]int{}
	for _, bedType := range bedTypeList {
		lowerText := strings.ToLower(bedType["en"])
		// has ["&"","and"]
		var pos int
		pos1 := strings.IndexByte(lowerText, "&"[0])
		if pos1 != -1 {
			pos = pos1
		}
		pos2 := strings.Index(lowerText, "and")
		if pos2 != -1 {
			pos = pos2
		}

		if pos != -1 {
			text1 := strings.TrimSpace(lowerText[:pos])
			text1Pos := strings.IndexByte(text1, " "[0])
			if text1Pos != -1 {
				quantity, err := strconv.Atoi(text1[:text1Pos])
				if err != nil {
					slog.Warn("WARN")
					continue
				}

				if _, ok := bedTypesMap[text1[text1Pos:]]; ok {
					quantityMapByBedType[text1[text1Pos:]] += quantity
				}
			}

			text2 := strings.TrimSpace(lowerText[pos:])
			text2Pos := strings.IndexByte(text2, " "[0])
			if text2Pos != -1 {
				quantity, err := strconv.Atoi(text2[:text2Pos])
				if err != nil {
					slog.Warn("WARN")
					continue
				}

				if _, ok := bedTypesMap[text2[text2Pos:]]; ok {
					quantityMapByBedType[text2[text2Pos:]] += quantity
				}
			}
		}

		pos = strings.IndexByte(lowerText, " "[0])
		if pos != -1 {
			if _, ok := bedTypesMap[strings.TrimSpace(lowerText[:pos])]; ok {
				quantityMapByBedType[strings.TrimSpace(lowerText[:pos])] += 1
			}

			quantity, err := strconv.Atoi(strings.TrimSpace(lowerText[:pos]))
			if err != nil {
				slog.Warn("WARN")
				continue
			}

			if _, ok := bedTypesMap[strings.TrimSpace(lowerText[pos+1:])]; ok {
				quantityMapByBedType[strings.TrimSpace(lowerText[pos+1:])] += quantity
			}
		}

		if _, ok := bedTypesMap[strings.TrimSpace(lowerText)]; ok {
			quantityMapByBedType[strings.TrimSpace(lowerText)] += 1
		}
	}

	var bedTypes []defaultBedType
	for bedType, quantity := range quantityMapByBedType {
		bedTypes = append(bedTypes, defaultBedType{
			BedTypeCode: bedTypesMap[bedType],
			Quantity:    quantity,
		})
	}

	return &defaultRoomTPAExtension{
		BedTypes: defaultBedTypes{
			BedType: bedTypes,
		},
	}
}
