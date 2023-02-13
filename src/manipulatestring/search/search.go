package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Contains("Hello, world!", "wo"))
	fmt.Println(strings.ContainsAny("Hello, world!", "w o"))
	fmt.Println(strings.Contains("Hello, world!", "ow"))
	fmt.Println(strings.ContainsAny("Hello, world!", "ow"))

	fmt.Println(strings.Count("Hello Helium", "He"))

	var r rune
	r = '하'
	fmt.Println(strings.ContainsRune("안녕하세요", r))
	fmt.Println(strings.HasPrefix("안녕하세요", "안"))
	fmt.Println(strings.HasSuffix("안녕하세요", "요"))

	fmt.Println(strings.Index("Hello, world!", "He"))
	fmt.Println(strings.Index("Hello, world!", "wor"))
	fmt.Println(strings.Index("Hello, world!", "ow"))

	fmt.Println(strings.IndexAny("Hello, world!", "eo"))
	fmt.Println(strings.IndexAny("Hello, world!", "f"))

	var c byte
	c = 'd'
	fmt.Println(strings.IndexByte("Hello, world!", c))
	c = 'f'
	fmt.Println(strings.IndexByte("Hello, world!", c))

	r = '집'
	fmt.Println(strings.IndexRune("우리 집", r))

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r)
	}

	fmt.Println(strings.IndexFunc("Go 언어", f))
	fmt.Println(strings.IndexFunc("Go Lang", f))

	fmt.Println(strings.LastIndex("Hello Hello Hello, world!", "Hello"))
	fmt.Println(strings.LastIndexAny("Hello, world!", "ol"))
	fmt.Println(strings.LastIndexFunc("Go 언어 안녕", f))
}
