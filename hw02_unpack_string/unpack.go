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
	var runes []rune = []rune(s)
	var first rune = runes[0]
	//var second rune = runes[1]

	if ok := unicode.IsLetter(first); !ok {
		return "", errors.New("not correct string")
	}

	var newString string = ""

	for i, r := range s {
		var okFirst bool
		var okSecond bool

		okFirst = unicode.IsLetter(rune(s[i]))

		if i+1 >= len(s) {
			fmt.Printf("последний элемент массива %s, записываем его\n", string(s[i]))
			newString += string(s[i])
			continue
		}
		okSecond = unicode.IsLetter(rune(s[i+1]))

		fmt.Println("текущее значение: ", string(s[i]))

		if !okFirst && !okSecond {
			if unicode.IsControl(rune(s[i])) {
				if unicode.IsDigit(rune(s[i+1])) {

					count, err := strconv.Atoi(string(s[i+1]))
					if err != nil {
						return "", fmt.Errorf("ошибка %s", err)
					}
					fmt.Println("Получаем: ", strings.Repeat(string(s[i]), count))
					newString += strings.Repeat(string(s[i]), count)
					continue
				}
			}
			err := fmt.Sprintf(`Передана не корректная строка %s: <%s> => <%s>`, s, string(s[i]), string(s[i+1]))

			return "", errors.New(err)
		}

		if okFirst && !okSecond {
			if unicode.IsControl(rune(s[i+1])) {
				fmt.Println("Это управляющий символ, пропустить")
				newString += string(s[i])
				continue
			}
			fmt.Printf("Дложны <%s> продублировать <%s> раз\n", string(s[i]), string(s[i+1]))

			count, err := strconv.Atoi(string(s[i+1]))
			if err != nil {
				return "", fmt.Errorf("ошибка %s", err)
			}

			fmt.Println("Получаем: ", strings.Repeat(string(s[i]), count))
			newString += strings.Repeat(string(s[i]), count)
		}

		if !okFirst && okSecond {
			fmt.Println("Пропускаем итерацию так как текущее значение это число: ", string(s[i]))
		}

		if okFirst && okSecond {
			fmt.Printf("Первый символ <%s> и второрй символ <%s> это не цифры и первый символ это буква ее и пишем \n", string(s[i]), string(s[i+1]))
			newString += string(s[i])
		}

		//fmt.Println(rune(s[i]))
		// fmt.Printf("Позиция %d, %s\n", i, string(s[i]))

		// fmt.Println(first)
		// fmt.Println(second)

		// okFirst := unicode.IsLetter(first)
		// okSecond := unicode.IsLetter(second)

		// if !okFirst {
		// 	first = second
		// 	second = r
		// 	continue
		// }

		// if okFirst && !okSecond {
		// 	count, err := strconv.Atoi(string(second))
		// 	if err != nil {
		// 		return "", fmt.Errorf("not correct string: %s", err)
		// 	}

		// 	n := strings.Repeat(string(first), count)

		// 	newString += n

		// 	first = second
		// 	second = rune(s[i+1])

		// 	continue
		// }

		// if (okFirst && okSecond) || (!okFirst && okSecond) {
		// 	newString += string(r)

		// 	first = second
		// 	second = r

		// 	continue
		// }

		// if ok := unicode.IsLetter(r); !ok {
		// 	fmt.Println("число", string(r))

		// 	count, err := strconv.Atoi(string(r))
		// 	if err != nil {
		// 		return "", fmt.Errorf("not correct string: %s", err)
		// 	}

		// 	fmt.Println("Полчаю число для count: ", count)
		// 	new := strings.Repeat(string(first), count)
		// 	fmt.Println("И дополняю строку: ", new)
		// 	newString += string(new)
		// 	//newRuns = append(newRuns, rune(new[0]))
		// } else {
		// 	fmt.Println("не число", string(r))
		// 	newString += string(r)
		// 	//newRuns = append(newRuns, r)
		// }

		//fmt.Printf("RUNE %c => %t\n", first, unicode.IsLetter(first))
		//fmt.Printf("Rune at index %d: %c\n", i, r)
		first = r

	}

	return newString, nil
}
