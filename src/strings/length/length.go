package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s1))
}
