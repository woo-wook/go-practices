package main

import "fmt"

func main() {
	var a [5]int

	a[2] = 3

	fmt.Println(a)

	var b = [...]int{
		1, 2, 3, 4, 5, // 여러 줄로 나열 시 마지막에 콤마가 필수다.
	}

	fmt.Println(b)
}
