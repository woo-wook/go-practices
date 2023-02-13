package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := []string{"Hello", "World"}
	fmt.Println(strings.Join(s1, " "))

	s2 := strings.Split("Hello, World!", " ")
	fmt.Println(s2[1])

	s3 := strings.Fields("Hello, world!")
	fmt.Println(s3[1])

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}

	s4 := strings.FieldsFunc("Hello안녕Hi", f)
	fmt.Println(s4)

	fmt.Println(strings.Repeat("Hello,", 30))

	fmt.Println(strings.Replace("Hello, world!", "world", "go", 1))
	fmt.Println(strings.Replace("Hello Hello Hello", "llo", "go", 2))
	fmt.Println(strings.Replace("Hello Hello Hello", "llo", "go", -1))

	fmt.Println(strings.Trim("Hello,~~ World!~~", "~~"))
	fmt.Println(strings.Trim(" Hello, World! ", " "))

	fmt.Println(strings.TrimLeft(" Hello, World! ~~  ", " "))
	fmt.Println(strings.TrimRight(" Hello, World! ~~", "~~"))

	f2 := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}

	s5 := "안녕 Hello 고 Go 고"

	fmt.Println(strings.TrimFunc(s5, f2))
	fmt.Println(strings.TrimLeftFunc(s5, f2))
	fmt.Println(strings.TrimRightFunc(s5, f2))

	fmt.Println(strings.TrimSpace("         hh       "))
}
