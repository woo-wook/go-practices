package main

import "fmt"

func main() {
	var num1 complex64 = complex(1, 2)
	var num2 complex64 = 4.2342 + 2.3527i
	var num3 complex64 = 1.e+3i
	var num4 complex64 = 7.27151e-10i
	var num5 complex128 = 1 + 2i
	var num6 complex128 = 5.32521e-10 + .1234e+2i

	var r1 float32 = real(num1)
	var i1 float32 = imag(num1)

	fmt.Printf("%f, %f\n", r1, i1)
	fmt.Println(num2, num3, num4, num5, num6)
}
