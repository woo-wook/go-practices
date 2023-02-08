package main

import "fmt"

func main() {
	var s1 string = "한글"
	var s2 string = "Go"

	fmt.Println(s1 == s2)
	fmt.Println(s1 + s2)
	fmt.Println("안녕하세요" + s1)
	fmt.Println(s1[0])
	fmt.Println(s1[1])
	fmt.Printf("%c", []rune(s1)[0]) // 한글은 이렇게..
}
