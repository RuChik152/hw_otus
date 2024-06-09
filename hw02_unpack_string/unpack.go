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

	if len(s) == 0 {
		return "", nil
	}

	var runes []rune = []rune(s)
	var first rune = runes[0]

	if ok := unicode.IsLetter(first); !ok {
		//err := fmt.Sprintf("передана не корректная строка %s", s)
		return "", ErrInvalidString
	}

	var newString strings.Builder

	for i, _ := range s {
		var okFirst bool
		var okSecond bool

		okFirst = unicode.IsLetter(rune(s[i]))

		if i+1 >= len(s) {
			fmt.Printf("последний элемент массива %s, записываем его\n", string(s[i]))

			newString.WriteByte(s[i])

			continue
		}
		okSecond = unicode.IsLetter(rune(s[i+1]))

		fmt.Println("текущее значение: ", string(s[i]))

		if !okFirst && !okSecond {
			if unicode.IsControl(rune(s[i])) {
				if unicode.IsDigit(rune(s[i+1])) {

					count, err := strconv.Atoi(string(s[i+1]))
					if err != nil {
						return "", ErrInvalidString
					}
					fmt.Println("Получаем: ", strings.Repeat(string(s[i]), count))

					arrByte := []byte(strings.Repeat(string(s[i]), count))
					newString.Write(arrByte)

					continue
				}
			}
			//err := fmt.Sprintf("Передана не корректная строка %s: <%s> => <%s>", s, string(s[i]), string(s[i+1]))

			return "", ErrInvalidString
		}

		if okFirst && !okSecond {
			if unicode.IsControl(rune(s[i+1])) {
				fmt.Println("Это управляющий символ, пропустить")

				newString.WriteByte(s[i])

				continue
			}

			fmt.Printf("Дложны <%s> продублировать <%s> раз\n", string(s[i]), string(s[i+1]))

			count, err := strconv.Atoi(string(s[i+1]))
			if err != nil {
				return "", ErrInvalidString
			}

			if count == 0 {
				continue
			}

			fmt.Println("Получаем: ", strings.Repeat(string(s[i]), count))

			arrByte := []byte(strings.Repeat(string(s[i]), count))
			newString.Write(arrByte)

		}

		if !okFirst && okSecond {
			fmt.Println("Пропускаем итерацию так как текущее значение это число: ", string(s[i]))
		}

		if okFirst && okSecond {
			fmt.Printf("Первый символ <%s> и второрй символ <%s> это не цифры и первый символ это буква ее и пишем \n", string(s[i]), string(s[i+1]))

			newString.WriteByte(s[i])

		}

	}

	return newString.String(), nil
}
