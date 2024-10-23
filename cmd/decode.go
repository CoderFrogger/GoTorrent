package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeBencode(benString string, startPos int) (interface{}, int, error) {
	switch {
	case benString[startPos] == 'i':
		return decodeBenInt(benString, startPos)
	case benString[startPos] == 'l':
		return nil, 0, nil
	case unicode.IsDigit(rune(benString[startPos])):
		return nil, 0, nil
	case benString[startPos] == 'd':
		return nil, 0, nil
	default:
		return nil, 0, nil

	}
}

func decodeBenInt(benString string, startPos int) (interface{}, int, error) {
	benIntEnd := strings.Index(benString[startPos:], "e") + startPos
	var decodedInt int
	var err error
	nextElementIndex := benIntEnd + 1

	if benString[1] == '-' {
		decodedInt, err = strconv.Atoi(benString[startPos+2 : benIntEnd])
		if err != nil {
			fmt.Println("Error during int decode: ", err)
			return nil, startPos, err
		}
		decodedInt = -(decodedInt)
	} else {
		decodedInt, err = strconv.Atoi(benString[startPos+1 : benIntEnd])
		if err != nil {
			fmt.Println("Error during int decode: ", err)
			return nil, startPos, err
		}
	}

	return decodedInt, nextElementIndex, nil
}

func decodeList(benString string, startPos int) ([]interface{}, int, error) {
	nextElementIndex := startPos + 1
	var err error

	decodedList := make([]interface{}, 0, 4)

	for benString[nextElementIndex] != 'e' {
		var decodedElement interface{}

		decodedElement, nextElementIndex, err = decodeBencode(
			benString,
			nextElementIndex,
		)
		if err != nil {
			fmt.Println("Error during list decode: ", err)
			return nil, startPos, err
		}

		decodedList = append(decodedList, decodedElement)
	}
	if nextElementIndex != len(benString) {
		nextElementIndex++
	}
	return decodedList, nextElementIndex, nil
}

func decodeBenStr(benString string, startPos int) (interface{}, int, error) {
	firstColonIndex := strings.Index(benString[startPos:], ":") + startPos

	numberSize := len(benString[startPos:firstColonIndex])
	strLength, err := strconv.Atoi(benString[startPos:firstColonIndex])
	if err != nil {
		fmt.Println("Error during string decode: ", err)
		return "", startPos, err
	}

	nextElementIndex := startPos + strLength + numberSize + 1
	return benString[firstColonIndex+1 : firstColonIndex+1+strLength], nextElementIndex, nil
}

func decodeDictionary(
	benString string,
	startPos int,
) (map[string]interface{}, int, error) {
	nextElementIndex := startPos + 1
	var err error

	decodedDict := make(map[string]interface{})

	for benString[nextElementIndex] != 'e' {
		var decodedKey, decodedValue interface{}

		decodedKey, nextElementIndex, err = decodeBencode(
			benString,
			nextElementIndex,
		)
		if err != nil {
			fmt.Println("Error during dictionary decode: ", err)
			return nil, startPos, err
		}

		decodedKeyString, ok := decodedKey.(string)
		if !ok {
			return nil, startPos, fmt.Errorf(
				"dictionary key not a string: %q",
				decodedKeyString,
			)
		}

		decodedValue, nextElementIndex, err = decodeBencode(
			benString,
			nextElementIndex,
		)
		if err != nil {
			fmt.Println("Error during dictionary decode: ", err)
			return nil, startPos, err
		}

		decodedDict[decodedKeyString] = decodedValue
	}

	return decodedDict, nextElementIndex, nil
}
