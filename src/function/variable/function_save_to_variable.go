package main

import "fmt"

func main() {
	var f1 func(a int, b int) int = plus
	f2 := plus

	fmt.Println(f1(1, 2))
	fmt.Println(f2(3, 4))

	var calculator = []func(int, int) int{
		plus,
		minus,
	}

	fmt.Println(calculator[0](10, 20))
	fmt.Println(calculator[1](10, 20))

	calculatorMap := map[string]func(int, int) int{
		"plus":  plus,
		"minus": minus,
	}

	fmt.Println(calculatorMap["plus"](10, 20))
	fmt.Println(calculatorMap["minus"](10, 20))
}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}
