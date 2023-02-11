package main

import "fmt"

func main() {
	a := [...]int{1, 3, 6, 3, 2}
	b := a

	b[3] = 5

	fmt.Println(a)
	fmt.Println(b)
}
