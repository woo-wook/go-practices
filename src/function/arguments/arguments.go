package main

import "fmt"

func main() {
	fmt.Println(plus(20, 30))
	fmt.Println(minus(20, 30))
}

func plus(a int, b int) (r int) {
	r = a + b

	return
}

func minus(a int, b int) int {
	return a - b
}
