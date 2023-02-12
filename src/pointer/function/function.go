package main

import "fmt"

func main() {
	i := 30
	fmt.Println(i)
	hello(&i)
	fmt.Println(i)
}

func hello(n *int) {
	*n = 2
}
