package main

import "fmt"

func main() {
	sum, _ := plusAndMinus(30, 20)
	_, diff := plusAndMinus(20, 5)

	fmt.Println(sum)
	fmt.Println(diff)
}

func plusAndMinus(a int, b int) (int, int) {
	return a + b, a - b
}
