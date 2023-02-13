package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go hello()

	for i := 0; i < 100; i++ {
		go printNumber(i)
	}

	fmt.Scanln()
}

func hello() {
	fmt.Println("Hello!")
}

func printNumber(n int) {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r))
	fmt.Println(n)
}
