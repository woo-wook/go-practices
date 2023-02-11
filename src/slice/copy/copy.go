package main

import "fmt"

func main() {
	source := []int{1, 2, 3}
	dest := make([]int, len(source))

	copy(dest, source)

	dest[0] = 9

	fmt.Println(source)
	fmt.Println(dest)
}
