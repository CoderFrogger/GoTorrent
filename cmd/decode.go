package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	BenInt        = 'i'
	BenList       = 'l'
	BenDictionary = 'd'
	IntegerEnd    = "e"
	ListEnd       = 'e'
)

func DecodeBencode(benString string, startPos int) (interface{}, int, error) {
	switch {
	case benString[startPos] == BenInt:
		return DecodeBenInt(benString, startPos)
	case benString[startPos] == BenList:
		return DecodeBenList(benString, startPos)
	case unicode.IsDigit(rune(benString[startPos])):
		return DecodeBenString(benString, startPos)
	case benString[startPos] == BenDictionary:
		return DecodeBenDictionary(benString, startPos)
	default:
		return nil, 0, fmt.Errorf(
			"Unknown bencode case: %x\n",
			benString[startPos],
		)

	}
}

func DecodeBenInt(benString string, startPos int) (interface{}, int, error) {
	benIntEnd := strings.Index(benString[startPos:], IntegerEnd) + startPos
	var decodedInt int
	var err error
	nextElementIndex := benIntEnd + 1

	if benString[1] == '-' {
		decodedInt, err = strconv.Atoi(benString[startPos+2 : benIntEnd])
		if err != nil {
			return nil, startPos, fmt.Errorf("Strconv failed: %x\n", err)
		}
		decodedInt = -(decodedInt)
	} else {
		decodedInt, err = strconv.Atoi(benString[startPos+1 : benIntEnd])
		if err != nil {
			return nil, startPos, fmt.Errorf("Strconv failed: %x\n", err)
		}
	}

	return decodedInt, nextElementIndex, nil
}

func DecodeBenList(benString string, startPos int) ([]interface{}, int, error) {
	nextElementIndex := startPos + 1
	var err error

	decodedList := make([]interface{}, 0, 4)

	for benString[nextElementIndex] != ListEnd {
		var decodedElement interface{}

		decodedElement, nextElementIndex, err = DecodeBencode(
			benString,
			nextElementIndex,
		)
		if err != nil {
			return nil, startPos, fmt.Errorf(
				"List bencode decode failed: %x\n",
				err,
			)
		}

		decodedList = append(decodedList, decodedElement)
	}
	if nextElementIndex != len(benString) {
		nextElementIndex++
	}
	return decodedList, nextElementIndex, nil
}

func DecodeBenString(benString string, startPos int) (interface{}, int, error) {
	firstColonIndex := strings.Index(benString[startPos:], ":") + startPos

	numberSize := len(benString[startPos:firstColonIndex])
	strLength, err := strconv.Atoi(benString[startPos:firstColonIndex])
	if err != nil {
		return "", startPos, fmt.Errorf("Integer conversion failed: %x\n", err)
	}

	nextElementIndex := startPos + strLength + numberSize + 1
	return benString[firstColonIndex+1 : firstColonIndex+1+strLength], nextElementIndex, nil
}

func DecodeBenDictionary(
	benString string,
	startPos int,
) (map[string]interface{}, int, error) {
	nextElementIndex := startPos + 1
	var err error

	decodedDict := make(map[string]interface{})

	for benString[nextElementIndex] != ListEnd {
		var decodedKey, decodedValue interface{}

		decodedKey, nextElementIndex, err = DecodeBencode(
			benString,
			nextElementIndex,
		)
		if err != nil {
			return nil, startPos, fmt.Errorf(
				"Dictionary key decode failed: %x\n",
				err,
			)
		}

		decodedKeyString, ok := decodedKey.(string)
		if !ok {
			return nil, startPos, fmt.Errorf(
				"dictionary key not a string: %q",
				decodedKeyString,
			)
		}

		decodedValue, nextElementIndex, err = DecodeBencode(
			benString,
			nextElementIndex,
		)
		if err != nil {
			return nil, startPos, fmt.Errorf(
				"Dictionary value decode failed: %x\n",
				err,
			)
		}

		decodedDict[decodedKeyString] = decodedValue
	}

	return decodedDict, nextElementIndex, nil
}
