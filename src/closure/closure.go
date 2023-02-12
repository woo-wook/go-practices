package main

import "fmt"

func main() {
	f := calc()

	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
}

func calc() func(int) int {
	a, b, c := 3, 5, 0

	return func(x int) int {
		fmt.Println(c)
		c = x
		return a*x + b
	}
}
