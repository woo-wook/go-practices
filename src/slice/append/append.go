package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	a = append(a, 4, 5, 6)

	fmt.Println(a)

	b := []int{9, 10, 11}
	a = append(a, b...)

	fmt.Println(a)
}
