package main

import "fmt"

func main() {
	var b1 bool = true
	var b2 bool = false

	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(!true)
	fmt.Println(b1)
	fmt.Println(b2)

	var num1 int = 3
	var num2 int = 10

	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 != num2)
}
