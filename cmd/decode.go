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
