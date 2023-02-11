package main

import "fmt"

func main() {
	var a []int = make([]int, 5)

	fmt.Println(a)

	var b = []int{1, 2, 5, 7, 8}

	fmt.Println(b)

	var s = make([]int, 5, 10) // 길이, 용량 순

	fmt.Println(cap(s))
	fmt.Println(len(s))
	fmt.Println(s[6]) // runtime error (길이는 초과 X)
}
