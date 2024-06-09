package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var newString strings.Builder

	if len(s) == 0 {
		return "", nil
	}

	if ok := unicode.IsLetter(rune(s[0])); !ok {
		return "", ErrInvalidString
	}

	for i := range s {
		if i+1 >= len(s) {
			newString.WriteByte(s[i])
			continue
		}

		if unicode.IsLetter(rune(s[i])) || unicode.IsSpace(rune(s[i])) {
			if unicode.IsDigit(rune(s[i+1])) {
				count, err := strconv.Atoi(string(s[i+1]))
				if err != nil {
					return "", ErrInvalidString
				}
				newString.Write([]byte(strings.Repeat(string(s[i]), count)))
				continue
			}
			newString.WriteByte(s[i])
			continue
		}

		if unicode.IsDigit(rune(s[i])) {
			if unicode.IsDigit(rune(s[i+1])) {
				return "", ErrInvalidString
			}
			continue
		}
	}

	return newString.String(), nil
}
