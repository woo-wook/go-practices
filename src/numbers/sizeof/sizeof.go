package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int8 = 1
	var num2 int16 = 1

	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
}
