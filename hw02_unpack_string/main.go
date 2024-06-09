package main

import "fmt"

func main() {
	var str string = `qwe\\\3`
	if s, err := Unpack(str); err != nil {
		fmt.Printf("ERROR: %q\n STR: %s", err, s)
	} else {
		fmt.Printf("RESULT: %q\n", s)
	}
}
