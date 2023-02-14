package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("He", "Hello 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("H.", "Hi")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[a-zA-Z0-9]", "Hello, 100")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[a-zA-Z0-9]*", "")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[0-9]+-[0-9]", "0111-111")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[0-9]+-[0-9]+", "0111-")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[0-9]+/[0-9]*", "0111/")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[0-9]+\\+[0-9]*", "0111+")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[^A-Z]", "hello")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("[가-힣]+", "안녕하세요")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("홍[가-힣]+동", "홍길동")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("유한[우-욱]+", "유한욱")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("^Hello", "Hello, world!")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("^Hello", "Go Hello, world!")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("world!$", "Go Hello, world!")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "hello.SetValue(20)")
	fmt.Println(matched)

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "hello.SetValue(20).x")
	fmt.Println(matched)
}
