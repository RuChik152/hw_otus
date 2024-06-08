package main

import "fmt"

func main() {
	var str string = "d\n1abc"
	if s, err := Unpack(str); err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("RESULT: ", s)
	}

}
