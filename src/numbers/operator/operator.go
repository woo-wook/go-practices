package main

import "fmt"

func main() {
	var num1 uint8 = 3
	var num2 uint8 = 2

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)
	fmt.Println(num1 << num2)
	fmt.Println(num1 >> num2)
	fmt.Println(^num1)

	var num3 float32 = 2.2

	// fmt.Println(num1 + num3) // 컴파일 에러!
	fmt.Println(float32(num1) + num3)

	var num4 uint16 = 10
	var num5 uint32 = 80000

	fmt.Println(num4 + uint16(num5))
}
