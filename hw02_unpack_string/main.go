package main

import "fmt"

func main() {
	var str string = "d\n5abc"
	if s, err := Unpack(str); err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Printf("RESULT: %q\n", s)
	}

}
