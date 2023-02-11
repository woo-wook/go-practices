package main

import "fmt"

func main() {
	a := [...]int{32, 69, 300, 23, 11}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, value := range a {
		fmt.Println(i, value)
	}

	for _, value := range a {
		fmt.Println(value)
	}
}
