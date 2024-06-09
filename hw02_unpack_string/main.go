package main

import "fmt"

func main() {
	var str string = "a4bc2d5e"
	if s, err := Unpack(str); err != nil {
		fmt.Printf("ERROR: %q\n", err)
	} else {
		fmt.Printf("RESULT: %q\n", s)
	}
}
