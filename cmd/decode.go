package cmd

import "unicode"

func decodeBencode(benString string, startPos int) {
	switch {
	case benString[startPos] == 'i':
	case benString[startPos] == 'l':
	case unicode.IsDigit(rune(benString[startPos])):
	case benString[startPos] == 'd':
	default:
		return

	}
}
