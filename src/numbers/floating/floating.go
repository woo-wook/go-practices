package main

import "fmt"

func main() {
	var num1 float32 = 0.1
	var num2 float32 = .35
	var num3 float32 = 132.73287

	var num4 float32 = 1e7
	var num5 float64 = .12345e+2
	var num6 float64 = 5.32521e-10

	fmt.Printf("%f, %f, %f, %f, %e %e", num1, num2, num3, num4, num5, num6)
}
