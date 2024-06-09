package main

import "fmt"

func main() {
	var str string = ""
	if s, err := Unpack(str); err != nil {
		fmt.Printf("ERROR: %q\n", err)
	} else {
		fmt.Printf("RESULT: %q\n", s)
	}

}
