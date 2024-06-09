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
	runes := []rune(s)

	if len(s) == 0 {
		return "", nil
	}

	if ok := unicode.IsLetter(rune(s[0])); !ok {
		return "", ErrInvalidString
	}

	for i := 0; i < len(runes); i++ {
		if i+1 >= len(s) {
			newString.WriteByte(byte(runes[i]))
			continue
		}

		if unicode.IsLetter(runes[i]) || unicode.IsSpace(runes[i]) {
			if unicode.IsDigit(runes[i+1]) {
				count, err := strconv.Atoi(string(runes[i+1]))
				if err != nil {
					return "", ErrInvalidString
				}
				newString.Write([]byte(strings.Repeat(string(runes[i]), count)))
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

		if unicode.IsPrint(runes[i]) && string(runes[i]) == "\\" {
			if unicode.IsDigit(runes[i+1]) || string(runes[i+1]) == "\\" {
				if i+2 < len(runes) && unicode.IsDigit(runes[i+2]) {
					count, err := strconv.Atoi(string(runes[i+2]))
					if err != nil {
						return "", ErrInvalidString
					}
					newString.Write([]byte(strings.Repeat(string(runes[i+1]), count)))
					i++
					i++
					continue
				}
				newString.WriteByte(byte(runes[i+1]))
				i++
				if i+2 >= len(runes) {
					break
				}
				continue
			} else {
				return "", ErrInvalidString
			}
		}
	}
	return newString.String(), nil
}
