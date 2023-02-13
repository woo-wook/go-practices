package main

import "fmt"

func main() {
	c := make(chan int)

	go sum(10, 50, c)

	n := <-c
	fmt.Println(n)
}

func sum(a int, b int, c chan int) {
	c <- a + b
}
