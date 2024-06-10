package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var newString strings.Builder
	runes := []rune(s)

	if len(s) == 0 {
		return "", nil
	}

	if ok := unicode.IsLetter(rune(s[0])); !ok {
		return "", ErrInvalidString
	}

	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))

		if i+1 >= len(s) {
			newString.WriteByte(byte(runes[i]))
			continue
		}

		if unicode.IsLetter(runes[i]) || unicode.IsSpace(runes[i]) {
			if unicode.IsDigit(runes[i+1]) {
				if err := writeData(&runes, &newString, i, 1); err != nil {
					return "", err
				}
				continue
			}
			newString.WriteByte(byte(runes[i]))
			continue
		}

		if unicode.IsDigit(runes[i]) {
			if unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}
			continue
		}

		if unicode.IsPrint(runes[i]) && isBackSlash(runes[i]) {
			var err error
			i, err = checkEscChar(&runes, &newString, i)
			if err != nil {
				return "", err
			}
			continue
		}
	}
	return newString.String(), nil
}

func writeData(runes *[]rune, s *strings.Builder, i int, factor int) error {
	count, err := strconv.Atoi(string((*runes)[i+factor]))
	if err != nil {
		return ErrInvalidString
	}
	s.Write([]byte(strings.Repeat(string((*runes)[i+(factor-1)]), count)))
	return nil
}

func isBackSlash(r rune) bool {
	return string(r) == "\\"
}

func isValidLastNumber(runes *[]rune, i int) bool {
	return i+2 < len((*runes)) && unicode.IsDigit((*runes)[i+2])
}

func checkEscChar(runes *[]rune, s *strings.Builder, i int) (int, error) {
	if !(unicode.IsDigit((*runes)[i+1]) || isBackSlash((*runes)[i+1])) {
		return i, ErrInvalidString
	}

	if isValidLastNumber(runes, i) {
		if err := writeData(runes, s, i, 2); err != nil {
			return i, err
		}
		return i + 2, nil
	}
	s.WriteByte(byte((*runes)[i+1]))
	i++
	if i+2 >= len((*runes)) {
		return i, nil
	}
	return i, nil
}
