package main

import "fmt"

func main() {
	var numPtr *int
	fmt.Println(numPtr)

	numPtr = new(int)
	fmt.Println(numPtr)
	fmt.Println(*numPtr)

	*numPtr = 10
	fmt.Println(*numPtr)

	num := 1
	ptr := &num

	fmt.Println(ptr)
	fmt.Println(&num)
}
