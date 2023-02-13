package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string = "한"
	fmt.Println(len(s))

	var r rune = '한'
	fmt.Println(utf8.RuneLen(r))

	var s1 string = "안녕하세여?"
	fmt.Println(utf8.RuneCountInString(s1))

	b := []byte("안녕하세요")

	r, size := utf8.DecodeRune(b)
	fmt.Printf("%c, %d\n", r, size)

	r, size = utf8.DecodeRune(b[3:])
	fmt.Printf("%c, %d\n", r, size)

	r, size = utf8.DecodeLastRune(b)
	fmt.Printf("%c, %d\n", r, size)

	r, size = utf8.DecodeLastRune(b[:len(b)-3])
	fmt.Printf("%c, %d\n", r, size)

	s2 := "Hello, World!"

	fmt.Printf("%c\n", s2[0])
	fmt.Printf("%c\n", s2[len(s2)-1])

	s3 := "안녕하세요"
	fmt.Printf("%c\n", s3[0])
	fmt.Printf("%c\n", s3[len(s3)-1])

	r, _ = utf8.DecodeRuneInString(s3)
	fmt.Printf("%c\n", r)

	r, _ = utf8.DecodeLastRuneInString(s3)
	fmt.Printf("%c\n", r)

	b2 := []byte(s3)
	b3 := []byte{0xff, 0xf1, 0xc1}

	fmt.Println(utf8.Valid(b2))
	fmt.Println(utf8.Valid(b3))
}
