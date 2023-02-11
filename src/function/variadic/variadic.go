package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	a := []int{1, 2, 3, 4, 5, 6, 9}
	fmt.Println(sum(a...)) // 슬라이스를 가변인자로
}

func sum(n ...int) int {
	total := 0

	for _, value := range n {
		total += value
	}

	return total
}
