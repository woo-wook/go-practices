package main

import "fmt"

func main() {
	var num1 = 10
	var num2 float32 = 3.2
	var num3 complex64 = 3.5 + 8.1i

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	// 개행이 없음
	fmt.Print(num1)
	fmt.Print(num2)
	fmt.Print(num3)
	fmt.Println()

	fmt.Printf("%d %f %f\n", num1, num2, num3)
	fmt.Printf("%03d %03f %f", num1, num2, num3)
}
