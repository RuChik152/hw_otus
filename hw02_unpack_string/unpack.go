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

	if len(s) == 0 {
		return "", nil
	}

	if ok := unicode.IsLetter(rune(s[0])); !ok {
		fmt.Println("Ошибка_3")
		return "", ErrInvalidString
	}

	for i, _ := range s {
		var okFirst bool
		var okSecond bool

		okFirst = unicode.IsLetter(rune(s[i]))
		if i+1 >= len(s) {

			newString.WriteByte(s[i])
			continue
		}
		okSecond = unicode.IsLetter(rune(s[i+1]))

		if !okFirst && !okSecond {
			if unicode.IsSpace(rune(s[i])) {
				if unicode.IsDigit(rune(s[i+1])) {

					count, err := strconv.Atoi(string(s[i+1]))
					if err != nil {
						fmt.Println("Ошибка_1: ", err)
						return "", ErrInvalidString
					}

					newString.Write([]byte(strings.Repeat(string(s[i]), count)))
					continue
				}
			}

			fmt.Println("Ошибка_2")
			return newString.String(), ErrInvalidString
		}

		if okFirst && !okSecond {
			if unicode.IsSpace(rune(s[i+1])) {

				newString.WriteByte(s[i])
				continue
			}

			if unicode.IsPrint(rune(s[i+1])) && string(s[i+1]) == "\\" {
				newString.WriteByte(s[i])
				continue
			}

			count, err := strconv.Atoi(string(s[i+1]))
			if err != nil {
				fmt.Println("Ошибка_4: ", err)
				return "", ErrInvalidString
			}

			if count == 0 {
				continue
			}

			newString.Write([]byte(strings.Repeat(string(s[i]), count)))
			continue
		}

		if !okFirst && okSecond {

			continue
		}

		if okFirst && okSecond {

			newString.WriteByte(s[i])
			continue
		}

	}

	return newString.String(), nil
}
